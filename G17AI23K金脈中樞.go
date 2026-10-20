package main

import "fmt"

// G17AI23KRecord 是金脈中樞的白話資料紀錄。
// 原始數字使用 string 保存，避免貨幣與代幣因浮點數而失真。
type G17AI23KRecord struct {
	Project    string // 專案名稱
	Creator    string // 創建者聲明
	Label      string // 使用者自訂標籤
	Source     string // 數字從哪裡來
	SeenAt     string // 看到或輸入的時間
	Asset      string // 資產名稱
	Currency   string // 幣別
	InputValue string // 使用者輸入多少，就保存多少
	Visible    string // 畫面看到多少，就保存多少
	Verified   string // 已核對、使用者申報、待核實
	Usable     string // 可動用、不可動用或未知

	CashIn     string // 實際流入
	CashOut    string // 實際流出
	Done       string // 已完成的部分
	Expected   string // 原本應收到或應完成的部分
	Contribution string // 專案貢獻
	Donation     string // 捐款
}

// PrintPlainLanguage 用最簡單的方式輸出紀錄。
func (r G17AI23KRecord) PrintPlainLanguage() {
	fmt.Println("=== G17AI23K｜金脈中樞｜白話紀錄 ===")
	fmt.Printf("專案：%s\n", r.Project)
	fmt.Printf("創建者聲明：%s\n", r.Creator)
	fmt.Printf("自訂標籤：%s\n", r.Label)
	fmt.Printf("資料來源：%s\n", r.Source)
	fmt.Printf("資料時間：%s\n", r.SeenAt)
	fmt.Printf("資產：%s %s\n", r.Asset, r.Currency)
	fmt.Printf("我輸入的數字：%s\n", r.InputValue)
	fmt.Printf("畫面看到的數字：%s\n", r.Visible)
	fmt.Printf("核實狀態：%s\n", r.Verified)
	fmt.Printf("可動用狀態：%s\n", r.Usable)
	fmt.Printf("現金流入：%s\n", r.CashIn)
	fmt.Printf("現金流出：%s\n", r.CashOut)
	fmt.Printf("應收到或應完成：%s\n", r.Expected)
	fmt.Printf("已收到或已完成：%s\n", r.Done)
	fmt.Printf("專案貢獻：%s\n", r.Contribution)
	fmt.Printf("捐款：%s\n", r.Donation)
	fmt.Println()
	fmt.Println("白話結論：")
	fmt.Printf("你輸入多少，我就記錄多少：%s\n", r.InputValue)
	fmt.Printf("你看到多少，我就呈現多少：%s\n", r.Visible)
	fmt.Printf("目前狀態：%s；可動用狀態：%s\n", r.Verified, r.Usable)
	fmt.Println("資料來源與核實狀態會分開顯示，不自行補造數字。")
}

func main() {
	record := G17AI23KRecord{
		Project:      "Project 2346Gm 金脈中樞",
		Creator:      "Meng-Hao Bu（13G17）",
		Label:        "我的資產紀錄",
		Source:       "使用者輸入／官方平台畫面／公開鏈上資料",
		SeenAt:       "由使用者填入",
		Asset:        "由使用者填入",
		Currency:     "由使用者填入",
		InputValue:   "由使用者填入",
		Visible:      "由使用者填入",
		Verified:     "使用者申報／待核實",
		Usable:       "尚未確認",
		CashIn:       "由使用者填入",
		CashOut:      "由使用者填入",
		Expected:     "由使用者填入",
		Done:         "由使用者填入",
		Contribution: "由使用者填入",
		Donation:     "由使用者填入",
	}

	record.PrintPlainLanguage()
}
