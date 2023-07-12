FROM registry.cn-zhangjiakou.aliyuncs.com/marksugar-k8s/golang:1.18.10-alpine3.17

RUN mkdir /test -p
WORKDIR /test
COPY . .

CMD ["go test"]