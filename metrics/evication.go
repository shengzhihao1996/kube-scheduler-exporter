package metrics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kube_scheduler_exporter/dindin"
	"net"
	"net/http"
	"sort"
	"strconv"
	"time"
)

func evication(hostname string) {
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

	resp, err := client.Get("http://10.68.65.50/api/v1/query?query=container_memory_rss")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	var dat map[string]interface{}
	json.Unmarshal(body, &dat)
	source := dat["data"].(map[string]interface{})["result"].([]interface{})

	var value = make(map[float64]string)
	var namespaces = make(map[string]string)

	for _, items := range source {
		cn := items.(map[string]interface{})["metric"].(map[string]interface{})["container_name"]
		vs := items.(map[string]interface{})["value"].([]interface{})[1].(string)
		hs := items.(map[string]interface{})["metric"].(map[string]interface{})["kubernetes_io_hostname"]
		if cn != "POD" && cn != nil && hs == hostname {
			v, _ := strconv.ParseFloat(vs, 64)
			pn := items.(map[string]interface{})["metric"].(map[string]interface{})["pod_name"].(string)
			ns := items.(map[string]interface{})["metric"].(map[string]interface{})["namespace"].(string)
			value[v] = pn
			namespaces[pn] = ns
		}
	}
	s1 := make([]float64, 0, len(value))
	for key, _ := range value {
		s1 = append(s1, key)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(s1)))
	for i := 0; i < 3; i++ {
		pod_name := value[s1[i]]
		go evicated(pod_name, namespaces[pod_name])
	}
	dindin.Evicatpod(hostname, value[s1[0]], value[s1[1]], value[s1[2]])
}

func evicated(pod_name string, namespaces string) {
	defer panics()
	export := shell("echo kubectl delete po  " + pod_name + " -n " + namespaces + "  --kubeconfig=/etc/config")
	fmt.Println(export)

}
