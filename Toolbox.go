// 對話資訊機器人
// 編譯方法：go build -buildmode=plugin -o Toolbox.so Toolbox.go

package main

import (
    tb "./TGBotLib"
    "fmt"
    "strings"
)

/* 相關字串 */
// modName: 模組名稱
const modName = "對話資訊"

// modDesc: 模組描述
const modDesc = "顯示對話的相關資訊。"

// settingsMsg: 因本模組沒有任何需設定項目，因此顯示的文字
const settingsMsg = "此模組沒有需要設定的項目。"

// msgSend: 詳細資訊訊息
const msgSend = ` [getUpdates 的詳細資訊訊息]

<此訊息的 update_id>: %d

[此訊息的相關資訊 (message)]
  <訊息識別碼 (message_id)>: %d
  <傳送日期 (unix) (date)>: %d
  <訊息文字 (text)>: %s

  [傳送者資訊 (from)]
    <使用者識別碼 (id)>: %d
    <使用者名字 (first_name)>: %s
    <使用者姓氏 (last_name)>: %s
    <使用者 ID (username)>: @%s
    <語言代碼 (language_code)>: %s

  [傳送之聊天室 (chat)]
    <聊天室識別碼 (id)>: %d
    <聊天室類型 (type)>: %s
    <聊天室標題 (title)>: %s
    <聊天室 ID (username)>: %s
    <使用者名字 (first_name)>: %s
    <使用者姓氏 (last_name)>: %s
`

// 資訊區塊
func Info() map[string]string {
    return map[string]string{
        "Name":        modName,
        "Author":      "pan93412",
        "Version":     "v1.0.0-20181226",
        "Description": modDesc,
    }
}

// 沒有設定。
func Settings() {
    fmt.Println(settingsMsg)
    return
}

// 處理區塊
func Handler(token string) {
    updates := tb.GetUpdatesBasic(token, true)
    if len(updates.Result) > 0 {
        msgdat := updates.Result[0]
        if strings.Contains(msgdat.Message.Text, "/getinfo") {
            // 啊哈哈…… Orz
            tb.SendMessageBasic(token, msgdat.Message.Chat.ID, fmt.Sprintf(msgSend,
                msgdat.UpdateID, msgdat.Message.MessageID, msgdat.Message.Date, msgdat.Message.Text,
                msgdat.Message.From.ID, msgdat.Message.From.FirstName, msgdat.Message.From.LastName,
                msgdat.Message.From.Username, msgdat.Message.From.LanguageCode,
                msgdat.Message.Chat.ID, msgdat.Message.Chat.Type, msgdat.Message.Chat.Title,
                msgdat.Message.Chat.Username, msgdat.Message.Chat.FirstName, msgdat.Message.Chat.LastName,
            ))
        }
    }
    return
}
