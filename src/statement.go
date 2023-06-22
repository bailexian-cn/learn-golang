package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main()  {
	var t uint64
	for {
		t++
		fmt.Println(fmt.Sprintf("%d time, %s", t, time.Now()))
		ok, body, err := checkTicket()
		fmt.Println(fmt.Sprintf("ticket: %v", ok))
		fmt.Println(fmt.Sprintf("err: %s", err))
		fmt.Println(fmt.Sprintf("err: %s", body))
		time.Sleep(3*time.Second)
	}
}

func checkTicket() (bool, string, error) {

	url := "https://www.allcpp.cn/allcpp/ticket/getTicketTypeList.do?eventMainId=1074"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-encoding", "gzip, deflate, br")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("cookie", "JSESSIONID=B4FFEB03780F4FF3D7455BF0ED39A2D4; JALKSJFJASKDFJKALSJDFLJSF=3871872364f20eb54670d42d687a61d78bd45a68f10.92.108.196_43129237187; Hm_lvt_75e110b2a3c6890a57de45bd2882ec7c=1682083681; token=OHRlhwk1Be0USIrui5MJ09jMDqHZMf2Ovr73vhMInv5Y2fZDgJrQw+4lQ+3cBNbPfh3E9rp6QiPjx6a+YyeL1Wgivl85pYFgb5ZtRiEg7svnB5Qs/sO0JVz4jq1iJs5UT0dIDFSaHZ5/tgwZp0O9lIEDak3Dum0quPIhf2yPSXM=; Hm_lpvt_75e110b2a3c6890a57de45bd2882ec7c=1682084191")
	req.Header.Add("origin", "https://cp.allcpp.cn")
	req.Header.Add("referer", "https://cp.allcpp.cn/")
	req.Header.Add("sec-ch-ua", `Chromium";v="112", "Google Chrome";v="112", "Not:A-Brand";v="99"`)
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", `"Windows"`)
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	jo := map[string]interface{}{}
	err := json.Unmarshal(body, &jo)
	if err != nil {
		return false, "", err
	}
	ticketTypeList := jo["ticketTypeList"].([]interface{})
	for _, t := range ticketTypeList {
		ticket := t.(map[string]interface{})
		remainderNum := ticket["remainderNum"]
		if remainderNum.(float64) > 0 {
			fmt.Println("true")
			return true, string(body), nil
		}
	}
	return false, "", nil
}