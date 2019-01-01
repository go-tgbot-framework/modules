# Telegram Bot Framework - 官方製作模組

## 模組列表
- CelebrateYear.go

  跨年機器人。您能使用此機器人在 2020, 2021 的元旦
  發出來自您機器人的跨年訊息！

- GGBot.go

  咕咕機器人。對機器人傳送「咕」或「g」就會傳送
  「咕咕咕～」配 GIF 動圖的訊息。

- Toolbox.go

  算是最有用的模組了吧（誤）
  可以傳送此聊天室的相關資訊的一個模組。

## 如何使用
1. 先編譯模組<br></br>
   `go build -buildmode=plugin -o (你想要的名稱).so (模組檔案)`<br></br>
   <br></br>
   (註) 目前不支援 Windows，若使用 Windows 10 可以使用 Bash 子系統，
        若使用 Windows 7 則可使用 cygwin 或直接安裝個虛擬機（文字介面即可）。

2. 放到 Telegram Bot Framework 的「modules」資料夾。<br></br>
   `cp (你上方取的名稱).so (Telegram Bot Framework 放置的位置)/modules`<br></br>
    <br></br>
   (註) 如果沒有 modules 資料夾，請先執行一次 Telegram Bot Framework
        即會自動產生。

3. 完成。
