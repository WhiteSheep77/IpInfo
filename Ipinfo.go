package Ipinfoby77

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// IPInfo 結構對應 API 回傳的 JSON 格式
type IPInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"` // 經緯度，例如 "37.3860,-122.0838"
	Org      string `json:"org"`
	Timezone string `json:"timezone"`
}

func IptoArea(ip string, token string) (res IPInfo, err error) {
	url := fmt.Sprintf("https://ipinfo.io/%s?token=%s", ip, token)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	} else {
		defer resp.Body.Close()
	}

	// 解碼 JSON 結果
	var info IPInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}

	// 顯示結果
	fmt.Printf("IP: %s\n", info.IP)
	fmt.Printf("City: %s\n", info.City)
	fmt.Printf("Region: %s\n", info.Region)
	fmt.Printf("Country: %s\n", info.Country)
	fmt.Printf("Location: %s\n", info.Loc)
	fmt.Printf("Org: %s\n", info.Org)
	fmt.Printf("Timezone: %s\n", info.Timezone)

	return
}
