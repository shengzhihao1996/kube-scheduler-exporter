package dindin

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
)

func Cordon(nodename string, v float64, export string) {
	value := float32(v) / 1024 / 1024 / 1000
	formt := `
	{
		"msgtype": "markdown",
		"markdown": {
			"title":"aiops",
			"text": "#### aiops \n #### NODE %s node_memory_Available is below 3G, the memory_Available is %.2fG. \n #### SchedulingDisabled %s now! \n #### output: %s" 

		}
	}`
	body := fmt.Sprintf(formt, nodename, value, nodename, export)
	jsonValue := []byte(body)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post("https://oapi.dingtalk.com/robot/send?access_token=b2637e4dd5eee826a7a29868cf0ff9e0818a40175dadc0840455bb0d0a03a4ab", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
}

func Uncordon(nodename string, v float64, export string) {
	value := float32(v) / 1024 / 1024 / 1000
	formt := `
	{
		"msgtype": "markdown",
		"markdown": {
			"title":"aiops",
			"text": "#### aiops \n #### NODE %s node_memory_Available is above 3G, the memory_Available is %.2fG. \n #### SchedulingAbled %s now! \n #### output: %s" 

		}
	}`
	body := fmt.Sprintf(formt, nodename, value, nodename, export)
	jsonValue := []byte(body)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post("https://oapi.dingtalk.com/robot/send?access_token=b2637e4dd5eee826a7a29868cf0ff9e0818a40175dadc0840455bb0d0a03a4ab", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
}

func Worring(x interface{}) {
	formt := `
	{
		"msgtype": "markdown",
		"markdown": {
			"title":"aiops",
			"text": "#### aiops \n #### 老子都挂了，赶紧看看咋回事了！！！ \n caught panic: %v" 

		}
	}`
	body := fmt.Sprintf(formt, x)
	jsonValue := []byte(body)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post("https://oapi.dingtalk.com/robot/send?access_token=b2637e4dd5eee826a7a29868cf0ff9e0818a40175dadc0840455bb0d0a03a4ab", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
}

func Evicatpod(hostname string, pn1 string, pn2 string, pn3 string) {
	formt := `
	{
		"msgtype": "markdown",
		"markdown": {
			"title":"aiops",
			"text": "#### aiops \n #### On node %s \n #### The following pod will be deleted: \n #### %s \n  #### %s \n #### %s  " 

		}
	}`
	body := fmt.Sprintf(formt, hostname, pn1, pn2, pn3)
	jsonValue := []byte(body)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post("https://oapi.dingtalk.com/robot/send?access_token=b2637e4dd5eee826a7a29868cf0ff9e0818a40175dadc0840455bb0d0a03a4ab", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
}
