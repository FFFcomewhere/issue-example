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
	authorization := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODUzMTk4NzksImlhdCI6MTY1Mzc4Mzg3OSwiand0VXNlcklkIjoxfQ.wpb9dqAQC7meyJ21_nrCSx7Q_30D28SRX0NAQ9STohc"
	contentType := "application/json"
	jsonStr := []byte(`{
						"issueid" : 1,
						"reName" : "newissue1",
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
						"reName" : "newmilestone2",
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

//----------------对tag模块的测试----------------------//
//测试 新建tag
func TestNewtag(t *testing.T) {
	url := "http://127.0.0.1:42002/tag/new"
	contentType := "application/json"
	jsonStr := []byte(`{
						"name":"tag3"}`)

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

//测试 按页查看tag
func TestGettagByPage(t *testing.T) {
	url := "http://127.0.0.1:42002/tag"
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

//测试 对单个tag进行操作
func TestShowSignaltag(t *testing.T) {
	url := "http://127.0.0.1:42002/tag/signal"
	//密钥，　经过登录或者注册可以获得，有过期时间，请自行登录后填写
	authorization := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODUwMjUyMDQsImlhdCI6MTY1MzQ4OTIwNCwiand0VXNlcklkIjo0fQ.zh0eUzFL7GFm0OAXiL6xbCKq_d3vKxLhS_e2IcNrVpw"
	contentType := "application/json"
	jsonStr := []byte(`{
						"name" : "tag3",
						"reName" : "newtag2",
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

//----------------对other模块的测试----------------------//
//测试 根据关键词查询　issues 比如issue.name,tag.name,comment.content
func TestIssueSearch(t *testing.T) {
	url := "http://127.0.0.1:42002/issue/search"
	//密钥，　经过登录或者注册可以获得，有过期时间，请自行登录后填写
	authorization := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODUwMjUyMDQsImlhdCI6MTY1MzQ4OTIwNCwiand0VXNlcklkIjo0fQ.zh0eUzFL7GFm0OAXiL6xbCKq_d3vKxLhS_e2IcNrVpw"
	contentType := "application/json"
	jsonStr := []byte(`{
						"keyword":"快乐"}`)

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
