package metrics

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func Metrics() {
	defer panics()
	for {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		reqest, err := http.NewRequest("GET", "https://10.68.0.1/api/v1/nodes", nil)
		reqest.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJzemgiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlY3JldC5uYW1lIjoiYWRtaW4tdG9rZW4teHE5Z2YiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiYWRtaW4iLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiI4M2VhMjA5NS05Njg2LTExZTktODk4My0yMDY3N2NkNGM5ZDgiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6c3poOmFkbWluIn0.lgRIrFk5d-znl1qtaUqupIro9iM23VsgX4PCKZuc6MZuEKRH1jWbsF56CcGmoaJ5Wr86oB2TvlHCdxhIU-fcli23XLCN4K5kNjaZG3JSu_0yy8fi5wZo5a3mDXCvoRTjgwqgzqCrGkMaacQ41M1k-cfRPNsh087wyXF-w9iKg6zHgSgIaMQf6HCdova_DP4f7y_ZW1sXfQKV-wFbzQr5jSWRaC6s3VPFpAZGQWbTB-lOBpb_DdX8wm2CrMyl2LBP_yejZ7IDP69FNGep7Dxre5h1N584QUrrmQdgHxx2jBKZk-U3tqLHDb5plmaoZJuW63rU5v_OPbBdZ55oNCpXqQ")
		resp, err := client.Do(reqest)
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		var dat map[string]interface{}
		json.Unmarshal(body, &dat)
		source := dat["items"].([]interface{})
		for _, items := range source {
			nodename := items.(map[string]interface{})["metadata"].(map[string]interface{})["name"].(string)
			role := items.(map[string]interface{})["metadata"].(map[string]interface{})["labels"].(map[string]interface{})["kubernetes.io/role"].(string)
			unschedulable := items.(map[string]interface{})["spec"].(map[string]interface{})["unschedulable"]
			block := items.(map[string]interface{})["metadata"].(map[string]interface{})["labels"].(map[string]interface{})["block"]
			maintain := items.(map[string]interface{})["metadata"].(map[string]interface{})["labels"].(map[string]interface{})["maintain"]
			evic := items.(map[string]interface{})["metadata"].(map[string]interface{})["labels"].(map[string]interface{})["evicated"]

			//可调度
			if role != "master" && unschedulable == nil && block == nil && maintain == nil {
				memorysearch(nodename, true, "false")
			}
			//不可调度
			if role != "master" && unschedulable != nil && block == nil && maintain == nil {

				memorysearch(nodename, false, "false")
			}
			//容器驱逐
			if role != "master" && unschedulable != nil && block != nil && maintain == nil && evic != nil {
				evic := items.(map[string]interface{})["metadata"].(map[string]interface{})["labels"].(map[string]interface{})["evicated"].(string)
				memorysearch(nodename, false, evic)
			}
		}
		time.Sleep(10 * time.Second)
	}

}
