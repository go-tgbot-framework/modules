// 跨年機器人
// 編譯方法：go build -buildmode=plugin -o CelebrateYear.so CelebrateYear.go

package main

import (
    tb "./TGBotLib"
    "fmt"
    "time"
    "io/ioutil"
    "strconv"
)

/* 相關常數 */
// 跨年機器人的年份設定位置
const yearPath = "CelebrateYear.yearoptions"

// 跨年機器人的聊天室 ID 設定位置
const chatIDPath = "CelebrateYear.chatidoptions"

// 資訊區塊
func Info() map[string]string {
    return map[string]string{
        "Name":        modName,
        "Author":      "pan93412",
        "Version":     "1.0",
        "Description": modDescription,
    }
}

/* 相關字串 */
// modName: 模組名稱
const modName = "跨年機器人"

// modDescription: 模組描述
const modDescription = "偵測是否到指定年份，若無則繼續倒數。"

// setYearPrompt: 設定年份時顯示的引導文字。
const setYearPrompt = "請輸入年份 (目前=%d，按 Enter 表示維持目前值)："

// setChatIDPrompt: 設定聊天室 ID 時顯示的引導文字。
const setChatIDPrompt = "TIPS: 使用 toolbox 模組取得群組的所有資訊\n請輸入聊天室 ID (目前=%d，按 Enter 表示維持目前值)："

// restartBotToApplyChanges: 設定完成後引導使用者重新啟動機器人框架的文字。
const restartBotToApplyChanges = "請重新啟動程式套用設定！"

// remainingTime: 離跨年剩餘時間。
// 請不要移除後方空白！
const remainingTime = "\r剩餘時間：%s    "

// timeOut: 時間到時顯示的文字。
const timeOut = "時間到！！！！"

// sendToGroup: 傳送到聊天室的訊息。
// %d (1, 2): 年份
// %f (3): 延遲秒數 (本機端延遲，不計傳送到伺服器的延遲。)
const sendToGroup = "%d 到了！(此訊息傳送時間已離 %d 整過 %f 秒)"

// err_optionsInvaild: 當傳入的 options 參數無效時。
const err_optionsInvaild = "傳入的 options 參數無效。"

/* 相關變數 */
// LoadSettings() 即為年份。
var year = LoadSettings("year").(int)
// year年 1月 1日 0時 0分 0秒 (本機時間)
var theTime = time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)

// 聊天室 ID
var chatID = LoadSettings("chatID").(int)

// LoadSettings: 載入設定
// 
// 若 options 為 "year"，回傳年份 (int)。
//
// 若 options 為 "chatid"，回傳聊天室 ID (int)。
//
// 讀取方法：型態斷言 (LoadSettings().(型態))
func LoadSettings(options string) interface{} {
    var theOptionsToRead string
    var isInt bool
    switch options {        
        case "year":
            theOptionsToRead = yearPath
            isInt = true
        case "chatid":
            theOptionsToRead = chatIDPath
            isInt = true
        default:
            panic(err_optionsInvaild)
    }
    
    theOption, err := ioutil.ReadFile(theOptionsToRead)
    if err != nil {
        currentSettings = []byte("")
    }
    
    if isInt {
        year, errAtoi := strconv.Atoi(string(currentSettings))
        if errAtoi != nil {
            year = 0
            return year
        }
    } else {
        return string(currentSettings)
    }
    
    return ""
}

// 設定區塊
func Settings() {
    var usryear string
    var usrchatid string
    
    // 年份讀取
    fmt.Printf(setYearPrompt, year)
    fmt.Scanln(&usryear)
    
    fmt.Printf(setYearPrompt, chatID)
    fmt.Scanln(&usrchatid)
    
    if year == "" {
        return
    } else {
        ioutil.WriteFile(settingsPath, []byte(year), 0644)
    }
    fmt.Println(restartBotToApplyChanges)
}

// 處理區塊
func Handler(token string) {
    nowTime := time.Now()
    fmt.Printf(remainingTime, theTime.Sub(nowTime).String())
    // 當 nowTime 超過 theTime，且當延遲秒數小於 0，大於 -1 時。
    if theTime.Before(nowTime) && theTime.Sub(nowTime).Seconds() < 0 && theTime.Sub(nowTime).Seconds() > -1 {
        fmt.Printf(timeOut)
        tb.SendMessageBasic(token, chatID, fmt.Sprintf(sendToGroup, year, year, nowTime.Sub(theTime).Seconds()))
    }
    return
}
