package process

import (
	"fmt"
	"encoding/json"
	"application/common/message"
)
func outputGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		return
	}
	
	// 显示内容
	content := fmt.Sprintf("用户id:\t%d 对大家说: \t%s", smsMes.UserId, smsMes.Content)

	fmt.Println(content)
	fmt.Println()
}

