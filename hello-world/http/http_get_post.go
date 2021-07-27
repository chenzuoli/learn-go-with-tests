package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//简单的http请求直接用http.Get和http.Post。需要操作headers，cookies或使用长连接和连接池，使用http.Client

// http get请求
func httpget() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// http post请求
func httppost() {
	resp, err := http.Post("http://www.baidu.com", "application/x-www-form-urlencode", strings.NewReader("name=abc"))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}

// http client
func httpclient() {
	client := &http.Client{}
	request, e := http.NewRequest("POST", "http://www.baidu.com", strings.NewReader("name=abc"))
	if e != nil {
		log.Fatal(e)
		return
	}
	request.Header.Set("Content_Type", "application/x-www-form-urlencoded")
	response, e := client.Do(request)
	if e != nil {
		log.Fatal(e)
		return
	}
	defer response.Body.Close()
	bytes, e := ioutil.ReadAll(response.Body)
	if e != nil {
		log.Fatal(e)
		return
	}
	fmt.Println(string(bytes))
}

func main() {
	//httpget()
	//httppost()
	httpclient()
}
