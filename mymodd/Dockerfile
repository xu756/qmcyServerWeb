FROM golang:alpine3.18

ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
ENV TZ=Asia/Shanghai

RUN go install github.com/cortesi/modd/cmd/modd@latest

CMD [ "modd" ]
