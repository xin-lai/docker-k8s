package main

import (
	"fmt"
	"os"
)
//环境变量
var envList = []string{
	//钉钉机器人地址
	"WEBHOOK",
	//@的手机号码
	"AT_MOBILES",
	//@所有人
	"IS_AT_ALL",
	//消息内容
	"MESSAGE",
	//消息类型（仅支持文本和markdown）
	"MSG_TYPE",
}

func main() {
	//获取环境变量
	envs := make(map[string]string)
	for _, envName := range envList {
		envs[envName] = os.Getenv(envName)
		//参数检查
		if envs[envName]=="" && envName != "AT_MOBILES" && envName != "IS_AT_ALL" {
			fmt.Println("envionment variable "+envName+" is required")
			os.Exit(1)
		}
	}

	if envs["AT_MOBILES"] == "" && envs["IS_AT_ALL"] == "" {
		fmt.Println("必须设置参数 AT_MOBILES 和 IS_AT_ALL 两者之一！")
		os.Exit(1)
	}

	builder, err := NewBuilder(envs)
	if err != nil {
		fmt.Println("BUILDER FAILED: ", err)
		os.Exit(1)
	}

	if err := builder.run(); err != nil {
		fmt.Println("BUILD FAILED", err)
		os.Exit(1)
	} else {
		fmt.Println("BUILD SUCCEED")
	}
}
