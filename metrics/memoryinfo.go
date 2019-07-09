package metrics

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

const (
	criticality float64 = 1000 * 1024 * 1024 * 3
)

func memorysearch(nodename string, schedu bool, evic string) {
	defer panics()
	client := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				client, err := net.DialTimeout(netw, addr, time.Second*1) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				client.SetDeadline(time.Now().Add(1 * time.Second)) //设置发送接收数据超时
				return client, nil
			},
			DisableKeepAlives: true,
		},
	}

	resp, err := client.Get("http://" + nodename + ":9100/metrics")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	components := strings.Split(*str, "\n")
	memoryinfo := func(components []string) float64 {
		var mem float64
		for _, v := range components {
			name := strings.Contains(v, "node_memory_MemAvailable_bytes")
			symbol := strings.Contains(v, "#")
			if name == true && symbol == false {
				components := strings.Split(v, " ")
				float, _ := strconv.ParseFloat(components[1], 64)
				mem = float
			}
		}
		return mem
	}(components)
	if schedu == true && memoryinfo < criticality {
		scheduler(nodename, memoryinfo, true)
	}
	if schedu == false && memoryinfo > criticality {
		scheduler(nodename, memoryinfo, false)
	}
	if schedu == false && memoryinfo < criticality && evic == "true" {
		evication(nodename)
	}
}
