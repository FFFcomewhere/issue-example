package issuecenter

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

//----------------对issue模块的测试----------------------//
//测试 新建issue
func TestNewIssue(t *testing.T) {
	url := "http://127.0.0.1:42002/issue/new"
	contentType := "application/json"
	jsonStr := []byte(`{
						"name":"issue_new1",
						"userName":"fzj",
						"tagName":"",
						"milestoneName":""}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
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

//测试 按页查看issue
func TestGetIssueByPage(t *testing.T) {
	url := "http://127.0.0.1:42002/issue"
	contentType := "application/json"
	jsonStr := []byte(`{
						"page":1,
						"pageSize":10}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
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

//测试 对单个issue进行操作
func TestShowSignalIssue(t *testing.T) {
	url := "http://127.0.0.1:42002/issue/signal"
	//密钥，　经过登录或者注册可以获得，有过期时间，请自行登录后填写
	authorization := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODUwMjUyMDQsImlhdCI6MTY1MzQ4OTIwNCwiand0VXNlcklkIjo0fQ.zh0eUzFL7GFm0OAXiL6xbCKq_d3vKxLhS_e2IcNrVpw"
	contentType := "application/json"
	jsonStr := []byte(`{
						"issueid" : 1,
						"reName" : "newissue2",
						"reTag" : "",
						"reMilestone" : "",
						"ifDelete" : false,	
						"addComment" : "",
						"deleteCommentid" : 0,
						"updateCommentid" : 0,
						"updateComment" : ""}`)

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

//----------------对milestone模块的测试----------------------//
//测试 新建milestone
func TestNewMilestone(t *testing.T) {
	url := "http://127.0.0.1:42002/milestone/new"
	contentType := "application/json"
	jsonStr := []byte(`{
						"name":"milestone3"}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
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

//测试 按页查看Milestone
func TestGetMilestoneByPage(t *testing.T) {
	url := "http://127.0.0.1:42002/milestone"
	contentType := "application/json"
	jsonStr := []byte(`{
						"page":1,
						"pageSize":10}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
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

//测试 对单个Milestone进行操作
func TestShowSignalMilestone(t *testing.T) {
	url := "http://127.0.0.1:42002/milestone/signal"
	//密钥，　经过登录或者注册可以获得，有过期时间，请自行登录后填写
	authorization := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODUwMjUyMDQsImlhdCI6MTY1MzQ4OTIwNCwiand0VXNlcklkIjo0fQ.zh0eUzFL7GFm0OAXiL6xbCKq_d3vKxLhS_e2IcNrVpw"
	contentType := "application/json"
	jsonStr := []byte(`{
						"milestoneid" : 1,
						"reName" : "newissue2",
						"ifDelete": false}`)

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
