package user

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

//测试用户注册
func TestRegister(t *testing.T) {
	url := "http://127.0.0.1:42001/user/register"
	contentType := "application/json"
	jsonStr := []byte(`{"username":"zkj", "password":"password"}`)

	resp, err := http.Post(url, contentType, bytes.NewBuffer(jsonStr))
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("handle error")
	}

	fmt.Println(string(body))
}

//测试用户登录
func TestLogin(t *testing.T) {
	url := "http://127.0.0.1:42001/user/login"
	contentType := "application/json"
	jsonStr := []byte(`{"username":"fzj", "password":"password"}`)

	resp, err := http.Post(url, contentType, bytes.NewBuffer(jsonStr))
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("handle error")
	}

	fmt.Println(string(body))
}

//测试获取用户信息
func TestGetUserInfo(t *testing.T) {
	url := "http://127.0.0.1:42001/user/info"
	contentType := "application/json"
	//密钥，　经过登录或者注册可以获得，有过期时间，请自行登录后填写
	authorization := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODUwMjUyMDQsImlhdCI6MTY1MzQ4OTIwNCwiand0VXNlcklkIjo0fQ.zh0eUzFL7GFm0OAXiL6xbCKq_d3vKxLhS_e2IcNrVpw"
	jsonStr := []byte(`{"userid":1,
						"mobile":"",
						"username":""}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Content-Type", contentType)

	client := http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("handle error")
	}

	fmt.Println(string(body))
}
