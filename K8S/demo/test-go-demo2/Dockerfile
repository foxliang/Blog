FROM golang:latest AS build
# 执行命令
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR /go/src/test_log
COPY . /go/src/test_log
RUN CGO_ENABLED=0 go build -v -o  main .

FROM alpine AS api
RUN mkdir /app
COPY --from=build /go/src/test_log /app
WORKDIR /app
ENTRYPOINT ["./main"]
