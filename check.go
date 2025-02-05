package rotateproxy

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type IPInfo struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func CheckProxyAlive(proxyURL string) (respBody string, timeout int64, avail bool) {
	proxy, _ := url.Parse(proxyURL)
	httpclient := &http.Client{
		Transport: &http.Transport{
			Proxy:           http.ProxyURL(proxy),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 20 * time.Second,
	}
	startTime := time.Now()
	resp, err := httpclient.Get("http://cip.cc/")
	if err != nil {
		return "", 0, false
	}
	defer resp.Body.Close()
	timeout = int64(time.Since(startTime))
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, false
	}
	if !strings.Contains(string(body), "地址") {
		return "", 0, false
	}
	return string(body), timeout, true
}

func StartCheckProxyAlive() {
	go func() {
		ticker := time.NewTicker(120 * time.Second)
		for {
			select {
			case <-crawlDone:
				fmt.Println("Checking")
				checkAlive()
				fmt.Println("Check done")
			case <-ticker.C:
				checkAlive()
			}
		}
	}()
}

func checkAlive() {
	proxies, err := QueryProxyURL()
	if err != nil {
		fmt.Printf("[!] query db error: %v\n", err)
	}
	for i := range proxies {
		proxy := proxies[i]
		go func() {
			respBody, timeout, avail := CheckProxyAlive(proxy.URL)
			if avail {
				fmt.Printf("%v 可用\n", proxy.URL)
				SetProxyURLAvail(proxy.URL, timeout, CanBypassGFW(respBody))
			} else {
				AddProxyURLRetry(proxy.URL)
			}
		}()
	}
}

func CanBypassGFW(respBody string) bool {
	return strings.Contains(respBody, "香港") ||
		strings.Contains(respBody, "台湾") ||
		strings.Contains(respBody, "澳门") || !strings.Contains(respBody, "中国")
}
