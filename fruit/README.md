# 往 mongo 中插入数据  
> db.fruit.insertMany([...])  
数据从 data.json 里 copy  

# proto  
退回到 codes 目录下，运行：  
> docker run --rm -v /Users/jsongo/code/docker/k8s/go-micro/codes:/go/src/ngbook go-alpine-builder go build main.go