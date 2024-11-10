package Ipinfoby77

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	LocX     float64
	LocY     float64
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

	res.LocX, res.LocY, _ = convertLocation(res.Loc)

	return res, string(bodyBytes), errUn

}

func convertLocation(location string) (float64, float64, error) {
	coords := strings.Split(location, ",")
	if len(coords) != 2 {
		return 0, 0, fmt.Errorf("invalid location format")
	}

	locx, err := strconv.ParseFloat(strings.TrimSpace(coords[0]), 64)
	if err != nil {
		return 0, 0, err
	}

	locy, err := strconv.ParseFloat(strings.TrimSpace(coords[1]), 64)
	if err != nil {
		return 0, 0, err
	}

	return locx, locy, nil
}
