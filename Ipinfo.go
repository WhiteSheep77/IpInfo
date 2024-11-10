package Ipinfoby77

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

func IptoArea(ip string, token string) (res IPInfo, ResbodyBytes string, reserr error) {
	url := fmt.Sprintf("https://ipinfo.io/%s?token=%s", ip, token)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
		return res, "", err
	} else {
		defer resp.Body.Close()
	}

	// 讀取 resp.Body 作為原始返回
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, string(bodyBytes), err
	}

	// 嘗試解析 JSON

	errUn := json.Unmarshal(bodyBytes, &res)

	if "" == res.IP && "" == res.City {
		return res, string(bodyBytes), errors.New(string(bodyBytes))
	}

	return res, string(bodyBytes), errUn

}
