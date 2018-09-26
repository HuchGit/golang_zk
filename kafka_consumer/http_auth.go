package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type machineTarge struct {
	Snname 		string `json:"snname"`
	Hostname 	string `json:"host_name"`
	Idc_room	string `json:"idc_room"`
	Cpu			string `json:"cpu"`
	Mem			string `json:"mem"`
	Disk        string `json:"disk"`
	Ip			string `json:"ip"`
	Ip_wan      string `json:"ip_wan"`
}

func http_info(url string, method string, body string, header map[string]string,key string,name string,sninfo string) (str_targe machineTarge){
	if url == "" || method == "" {
		return str_targe
	}
	url = url +"?"+"sn="+key+"&items="+sninfo
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	http_client := &http.Client{Timeout: 20 * time.Second, Transport: tr}
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return str_targe
	}
	if header != nil {
		for key, val := range header {
			req.Header.Add(key, val)
		}
	}
	// do a request
	resp, err := http_client.Do(req)
	if err != nil {
		return str_targe
	}
	defer resp.Body.Close()

	byte_body,err:=ioutil.ReadAll(resp.Body)
	fmt.Println(string(byte_body))
	var numberstr map[string]interface{}
	err = json.Unmarshal([]byte(byte_body),&numberstr)
//	for _,tmp := range numberstr{
		str_targe.Snname = key
		str_targe.Idc_room = name
		str_targe.Hostname = numberstr["hostname"].(string)
		str_targe.Ip = numberstr["ip"].(string)
		str_targe.Ip_wan = numberstr["ip_wan"].(string)
		str_targe.Cpu = numberstr["cpu"].(string)
		str_targe.Mem = numberstr["memory"].(string)
		str_targe.Disk = numberstr["disk"].(string)

//	}
	fmt.Println(str_targe)
	return str_targe
}

func http_do(url string, method string, body string, header map[string]string) (map[string]string, error) {
	// parameter check
	sn_map := make(map[string]string)
	if url == "" || method == "" {
		return sn_map, errors.New("param is invald")
	}
	// make the http client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	http_client := &http.Client{Timeout: 20 * time.Second, Transport: tr}
	// make the request
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	if header != nil {
		for key, val := range header {
			req.Header.Add(key, val)
		}
	}
	// do a request
	resp, err := http_client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	byte_body,err:=ioutil.ReadAll(resp.Body)
	fmt.Println(string(byte_body))
	var numberstr []map[string]interface{}
	err = json.Unmarshal([]byte(byte_body),&numberstr)
	var sn_list []string
	for _,tmp := range numberstr{
		sn:=tmp["sn"].(string)
		sn_map[sn]=tmp["idc_room__full_name"].(string)
	}
	fmt.Println(sn_list)
	return sn_map,err
}

func main() {

	auth_token := "Token d0da7315996d2bf19c85c9815609d7ae54ed7b07"
	url := "http://xxxx/api/rest/bus_host/?bus=public.cdn"
	url1 := "http://xxx/api/rest/host_items/"
	sninfo:="idc_room,hostname,ip,ip_wan,cpu,memory,disk"

	targe_map := []machineTarge{}
	header := map[string]string{"Authorization":auth_token,"Content-Type":"application/json"}
	result, err := http_do(url,"GET","",header)
	fmt.Printf("response:\n%v\nerror:\n%v\n", result, err)
	for key,name := range result{
		targe := http_info(url1,"GET","",header,key,name,sninfo)
		targe_map = append(targe_map,targe)
	}
	fmt.Println(targe_map)
	b, err := json.Marshal(targe_map)
	fmt.Println(string(b))
}
