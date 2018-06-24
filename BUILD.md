# 编译、部署一个微服务（以fruit举例）

## proto  
退回到 go-micro 目录下，运行：  
> ./gen-micro-pb.sh ngbook/fruit/proto/fruit.api.proto  

## 怎么构建可在 alpine 上运行的二进制文件？  
### 如果 docker 版本在17.05以上  
[参考这里](https://docs.docker.com/develop/develop-images/multistage-build/)  
使用 docker/17.05+ 里的 Dockerfile  

### 如果 docker 版本比较低  
> docker run --rm -it -v /tmp/bin:/go/bin jsongo/go-alpine-builder go get -u github.com/ngbook/rest-api/fruit  
结束后，可以在 /tmp/bin 里找到相关的可运行文件  
使用 docker/17.05- 里的 Dockerfile 进行新镜像的构建  

## 之后，把构建好的镜像推送到 docker hub 上使用  