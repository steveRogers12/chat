package sms

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type ResponseXML struct {
	XMLName         xml.Name `xml:"returnsms"`
	Returnstatus    string   `xml:"returnstatus"`
	Message         string   `xml:"message"`
	Remainpoint     int64   `xml:"remainpoint"`
	TaskID          string `xml:"taskID"`
	SuccessCounts   int `xml:"successCounts"`
}
var smsTitle = map[int]string{
	0:"【全村希望】",
}
//发送短信
func SendSms(text,phone string) (bool,error){
	client := &http.Client{}
	smsConfig := viper.GetStringMap("sms")
	smsUrl := smsConfig["url"].(string)
	fmt.Printf("smsUrl%v", smsUrl)

	pData := url.Values{}
	pData.Add("action", "send")
	pData.Add("userid", "")
	pData.Add("account", smsConfig["account"].(string))
	pData.Add("password", smsConfig["password"].(string))
	pData.Add("mobile", phone)
	pData.Add("content", smsTitle[0]+text)

	fmt.Printf("pData%v", pData)
	//提交请求
	reqest, err := http.NewRequest("POST", smsUrl, strings.NewReader(pData.Encode()))
	if err != nil {
		panic(err)
	}
	//增加header选项
	reqest.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	//处理返回结果
	response, err := client.Do(reqest)
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("error-client.Do:%v", err)
		return false, errors.New("发送失败")
	}
	respBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("error-ioutil.ReadAll:%v", err)
		return false, errors.New("发送失败")
	}
	fmt.Printf("respBytes%v", string(respBytes))
	bs := ResponseXML{}
	//把xml数据解析成对象
	err = xml.Unmarshal(respBytes, &bs)
	if err != nil {
		fmt.Printf("error-xml.Unmarshal:%v", err)
		return false, errors.New("xml解析失败")
	}
	if bs.SuccessCounts != 1 {
		fmt.Printf("error-发送失败:%v", bs.Message)
		return false, errors.New(bs.Message)
	}

	return true,nil
}