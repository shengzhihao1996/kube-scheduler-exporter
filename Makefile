install: go-build docker-build
go-build:
	mkdir -p bin
	cp `which kubectl` .
	cp /root/.kube/config .
	docker run  --rm -v `pwd`:/go/src registry.cn-huhehaote.aliyuncs.com/shengzhihao/go-build:v1  go build -o /go/bin/app /go/src/main.go
docker-build:
	docker build -t aiops:release-1.0.0 .
