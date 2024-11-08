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

func IptoArea(ip string, token string) (res IPInfo, reserr error) {
	url := fmt.Sprintf("https://ipinfo.io/%s?token=%s", ip, token)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
		return res, err
	} else {
		defer resp.Body.Close()
	}

	// 解碼 JSON 結果
	var info IPInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
		return res, err
	}

	return info, err
}
