package metrics

import (
	"fmt"
	"kube_scheduler_exporter/dindin"
	"log"
	"os/exec"
	"time"
)

func scheduler(nodename string, value float64, schedu bool) {
	defer panics()
	if schedu == true {
		export := shell(" kubectl cordon " + nodename + " --kubeconfig=/etc/config")
		_ = shell(" kubectl label node  " + nodename + " block=true  --kubeconfig=/etc/config")
		go timewait(nodename)
		fmt.Println(export)
		dindin.Cordon(nodename, value, export)
		fmt.Println(nodename, "禁止调度")
	}
	if schedu == false {
		export := shell(" kubectl uncordon " + nodename + " --kubeconfig=/etc/config")
		fmt.Println(export)
		dindin.Uncordon(nodename, value, export)
		fmt.Println(nodename, "解除限制")
	}
}

//阻塞20分钟
func timewait(nodename string) {
	time.Sleep(2 * time.Minute)
	_ = shell(" kubectl label node  " + nodename + " evicated=test  --kubeconfig=/etc/config")
	time.Sleep(18 * time.Minute)
	fmt.Println(shell(" kubectl label node  " + nodename + " block-  --kubeconfig=/etc/config"))
	_ = shell(" kubectl label node  " + nodename + " evicated-  --kubeconfig=/etc/config")

}

func panics() {
	if x := recover(); x != nil {
		log.Printf("caught panic: %v", x)
		dindin.Worring(x)
	}
}
func shell(s string) string {
	cmd := exec.Command("/bin/sh", "-c", s)
	output, _ := cmd.Output()
	return string(output)
}
