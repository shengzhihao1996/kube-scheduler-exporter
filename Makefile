install: go-build docker-build
go-build:
	mkdir -p bin
	cp `which kubectl` .
	cp /root/.kube/config .
	docker run  --rm -v `pwd`:/go/src/kube_scheduler_exporter registry.cn-huhehaote.aliyuncs.com/shengzhihao/go-build:v1  go build -o /go/src/kube_scheduler_exporter/main   /go/src/kube_scheduler_exporter/main.go
docker-build:
	docker build -t aiops:release-1.0.3 .
