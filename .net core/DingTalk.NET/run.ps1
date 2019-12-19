docker build --rm -f "Dockerfile" -t dingtalk.net:latest .

docker run --rm -e "WEBHOOK=https://oapi.dingtalk.com/robot/send?access_token=40eb535b2ff6baee6db2f4b76a1a35d5852a3ac9c2d8cfb8e4e7c35219476df2" `
    -e "MESSAGE=*使用.NET Core发送钉钉消息。*" `
    -e "IS_AT_ALL=true" `
    -e "MSG_TYPE=markdown" `
    -d dingtalk.net