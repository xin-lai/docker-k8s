package main
// 导入包
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// 结构
type Builder struct {
	Webhook        string
	AtMobiles      []string
	IsAtAll        bool
	Message        string
	payload interface{}
}

// 组装消息格式
func NewBuilder(envs map[string]string) (*Builder, error) {
	b := &Builder{}

	b.Webhook = envs["WEBHOOK"];

	if strings.ToLower(envs["IS_AT_ALL"]) == "true" {
		b.IsAtAll = true
	}

	b.AtMobiles = strings.Split(envs["AT_MOBILES"], ",")
	at := At{
		AtMobiles: b.AtMobiles,
		IsAtAll:   b.IsAtAll,
	}

	if envs["MESSAGE"] != "" {
		b.Message = envs["MESSAGE"]
		switch envs["MSG_TYPE"] {
			case "text":
				text := Text{
					Content: b.Message,
				}
				info := TextWebhook{
					Msgtype: "text",
					Text:    text,
					At:      at,
				}
				b.payload = info
				return b, nil
			case "markdown":
				md := Markdown{
						Title: "钉钉通知",
						Text:  b.Message,
				}
				b.payload = MarkdownWebHook{
					Msgtype:  "markdown",
					Markdown: md,
					At:       at,
				}
				return b, nil
		default:
			return nil, fmt.Errorf("不支持的消息类型！")
		}
	}
	return nil, fmt.Errorf("尚不支持其他格式！")
}

func (b *Builder) run() error {
	if err := b.callWebhook(); err != nil {
		return err
	}
	return nil
}

//调用钉钉webhook
func (b *Builder) callWebhook() error {
	payload, _ := json.Marshal(b.payload)
	fmt.Printf("sending webhook info: %s\n", string(payload))
	body := bytes.NewBuffer(payload)
	res, err := http.Post(b.Webhook, "application/json;charset=utf-8", body)
	if err != nil {
		return err
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	var resultJSON interface{}
	if err = json.Unmarshal(result, &resultJSON); err != nil {
		return err
	}

	fmt.Println(resultJSON)
	fmt.Println("Send webhook succeed.")
	return nil
}