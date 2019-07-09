## if you want to build image, please install docker, you must have access to an external network, then use: 

```
make install
```

###### 可调度，可用内存小于3G，非阻塞，非维护  ——>  禁止调度，添加阻塞标签，20m后去掉标签
###### 不可调度，可用内存大于3G，非阻塞，非维护  ——>  允许调度，解除限制
###### 节点维护标签： maintain=true 程序逻辑判断时会排除这些节点
###### 程序或许会强依赖集群角色标签，务必补充（kubernetes.io/role=master，kubernetes.io/role=node）
