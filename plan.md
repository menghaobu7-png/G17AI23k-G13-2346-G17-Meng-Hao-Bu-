# PROJECT 2346 · 自有資金事件鏈

## 實作方向
建立一個單頁、可直接開啟的繁體中文金融事件鏈介面，將需求中的「入金／出金」呈現為本地追加式帳本操作。前端以 localStorage 保存事件，使用 SHA-256 產生 previous_hash 與 event_hash，讓使用者可以查看鏈狀事件、餘額、節點狀態與驗證結果。這是本地展示，不連接銀行、交易所或真實支付系統。

## 主要功能
- 入金、出金表單與金額驗證。
- 每次操作追加一筆不可由 UI 編輯的事件。
- 本地 SHA-256 雜湊鏈驗證與事件明細。
- 帳戶餘額、累計入金、累計出金、事件數統計。
- 節點／共識狀態視覺化，明確標示 LOCAL DEMO。
- 重設展示資料，避免誤以為是真實帳務。

## 設計
候選風格：Industrial Cyber Ledger (0.06)、Quiet Banking Console (0.07)、Neon Audit Room (0.05)。選用 **Industrial Cyber Ledger**。

- **Design Movement**：賽博龐克工業資料終端，將粗獷金屬框架與精準帳務資訊並置。
- **Core Principles**：可核驗、低誤導、資訊優先、操作有回饋。
- **Color Philosophy**：深墨藍作為安全基底，青綠代表鏈完整與入金，珊瑚紅只用於出金與警示，紫色作為 PROJECT 2346 的識別光源。
- **Layout Paradigm**：左側固定導航＋右側縱向稽核流，不以傳統卡片牆作為唯一結構。
- **Signature Elements**：事件鏈時間線、掃描線背景、雜湊碼膠囊。
- **Interaction Philosophy**：每次操作都立即顯示事件 ID、鏈高度與新的帳戶餘額；不隱藏示範限制。
- **Animation**：新增事件時短暫高亮，鏈驗證按鈕以微弱掃描線回饋，不使用干擾性動畫。
- **Typography System**：Inter／Noto Sans TC 為主，等寬字體用於事件 ID、雜湊與數值。
- **Brand Essence**：給需要看懂資金事件證據的人，一個不假裝接入真金流的自有帳本介面；可靠、克制、技術感。
- **Brand Voice**：標題直接，說明不誇大。例：「每一筆變更，都留下證據。」、「本地展示，不代表銀行入帳。」
- **Wordmark & Logo**：23／46 雙層數字標記，像兩個相連的帳本區塊。
- **Signature Brand Color**：電光青綠 #42f5c5。

## 專案結構
- `index.html`：完整頁面、樣式與互動邏輯，無外部依賴。
- `public/manus-routes.json`：頁面路由清單。
- `TODO.md`：交付結果與驗收條件。
