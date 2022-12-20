package message

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-uuid"
	eventmethod "github.com/waro163/feishu_robot/event-method"
	"github.com/waro163/feishu_robot/feishu"
	msg "github.com/waro163/feishu_robot/feishu/message"
	"github.com/waro163/feishu_robot/gtp"
	"strings"
)

func init() {
	eventmethod.RegisterEventMethod("im.message.receive_v1", HandleMsgEvent)
	eventmethod.RegisterEventMethod("im.message.message_read_v1", HandleMsgReadEvent)
}

type Text struct {
	Text string `json:"text"`
}

func HandleMsgEvent(header map[string]string, event map[string]interface{}) error {
	message := event["message"]
	if message == nil {
		return nil
	}
	m := message.(map[string]interface{})
	content := m["content"].(string)
	messageId := m["message_id"].(string)
	if content == "" || messageId == "" {
		return nil
	}
	text := &Text{}
	err := json.Unmarshal(([]byte)(content), text)
	if err != nil {
		fmt.Printf("%v", err)
		return nil
	}
	t := text.Text
	t = t[strings.Index(t, " ")+1:]
	println("<<<<<<<<<<<<<<")
	println(t)
	completions, err := gtp.Completions(t)
	if err != nil {
		fmt.Printf("%v", err)
		return nil
	}
	body := make(map[string]interface{})
	//{
	//    "content": "{\"text\":\"<at user_id=\\\"ou_155184d1e73cbfb8973e5a9e698e74f2\\\">Tom </at> test content\"}",
	//    "msg_type": "text",
	//    "uuid": "a0d69e20-1dd1-458b-k525-dfeca4015204"
	//}

	cbs, err := json.Marshal(&Text{Text: completions})
	if err != nil {
		fmt.Printf("%v", err)
		return nil
	}

	body["content"] = strings.ReplaceAll(string(cbs), "\n", "\\\n")
	body["msg_type"] = "text"
	generateUUID, err := uuid.GenerateUUID()
	if err != nil {
		fmt.Printf("%v", err)
		return nil
	}
	body["uuid"] = generateUUID
	accessToken, err := feishu.GetTenantAccessToken()
	if err != nil {
		fmt.Printf("%v", err)
		return nil
	}
	_, err = msg.ReplyMsg(messageId, accessToken, body)
	if err != nil {
		fmt.Printf("%v", err)
		return nil
	}
	return nil
}

func HandleMsgReadEvent(header map[string]string, event map[string]interface{}) error {
	return nil
}
