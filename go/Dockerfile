FROM golang:1.10-alpine as builder

WORKDIR /go/src/component-dingding

COPY ./ /go/src/component-dingding

RUN set -ex && \
go build -v -o /go/bin/component-dingding \
-gcflags '-N -l' \
./*.go

FROM alpine
RUN apk update && apk add ca-certificates

COPY --from=builder /go/bin/component-dingding /usr/bin/
CMD ["component-dingding"]

# 注意不要单独使用 MAINTAINER 指令，MAINTAINER已被Label标签代替
LABEL MAINTAINER "xinlai@xin-lai.com"
# LABEL指令用于将元数据添加到镜像，支持键值对和JSON，我们可以使用 docker inspect 命令来查看
LABEL DingtalkComponent='{\
  "description": "使用钉钉发送通知消息.",\
  "input": [\
    {"name": "WEBHOOK", "desc": "必填, 钉钉机器人Webhook地址"},\
    {"name": "AT_MOBILES", "desc": "非必填，被@人的手机号"},\
    {"name": "IS_AT_ALL", "desc": "非必填，@所有人时:true, 否则为:false"},\
    {"name": "MESSAGE", "desc": "必填，自定义发送的消息内容"}，\
    {"name": "MSG_TYPE", "desc": "必填，自定义发送的消息类型，目前仅支持text和markdown"}\
  ]\
}'
