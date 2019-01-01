// 咕咕機器人
package main

import (
    tb "./TGBotLib"
    "strings"
)

var photourl = "https://i.imgur.com/begraED.png"
var sendmsg = "咕咕咕～"

// 資訊區塊
func Info() map[string]string {
    return map[string]string{
        "Name":        "咕咕機器人",
        "Author":      "pan93412",
        "Version":     "v1.0-20181226",
        "Description": "偵測使用者輸入的文字是否包含「咕」或「g」這一類文字。如果有則發出咕咕圖。",
    }
}

// 設定區塊
func Settings() {
    return
}

// 處理區塊
func Handler(token string) {
    updates := tb.GetUpdatesBasic(token, true)
    if len(updates.Result) > 0 {
        if strings.Contains(updates.Result[0].Message.Text, "g") || strings.Contains(updates.Result[0].Message.Text, "咕") {
            tb.SendMedia(token, "document", updates.Result[0].Message.Chat.ID, photourl, sendmsg, false, -1)
        }
    }
    return
}
