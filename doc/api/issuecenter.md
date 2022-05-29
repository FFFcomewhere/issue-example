
### 1. "按页展示提案"

1. 路由定义

- Url: /issue
- Method: POST
- Request: `IssueReq`
- Response: `IssueResp`

2. 请求定义


```golang
type IssueReq struct {
	Page int64 `json:"page"` //页码
	PageSize int64 `json:"pageSize"` //当前页面显示的issue数量
}
```


3. 返回定义


```golang
type IssueResp struct {
	List []IssueInfo `json:"list"`
}
```
  


### 2. "创建提案"

1. 路由定义

- Url: /issue/new
- Method: POST
- Request: `IssueNewReq`
- Response: `IssueNewResp`

2. 请求定义


```golang
type IssueNewReq struct {
	Name string `json:"name"`
	UserName string `json:"userName"`
	TagName string `json:"tagName"`
	MilestoneName string `json:"milestoneName"`
}
```


3. 返回定义


```golang
type IssueNewResp struct {
}
```
  


### 3. "操作单个提案,评论/设置标签或者里程碑"

1. 路由定义

- Url: /issue/signal
- Method: POST
- Request: `IssueSignalReq`
- Response: `IssueSignalResp`

2. 请求定义


```golang
type IssueSignalReq struct {
	Issueid int64 `json:"issueid"`
	ReName string `json:"reName"` //重命名标题
	ReTag string `json:"reTag"` //重命名标签
	ReMilestone string `json:"reMilestone"` //重命名里程碑
	IfDelete bool `json:"ifDelete"` //是否删除提案
	AddComment string `json:"addComment"` //添加评论
	DeleteCommentid int64 `json:"deleteCommentid"` //被删除评论的id
	UpdateCommentid int64 `json:"updateCommentid"` //被修改评论的id
	UpdateComment string `json:"updateComment"` //修改评论
}
```


3. 返回定义


```golang
type IssueSignalResp struct {
	Issue IssueInfo `json:"issue"` //提案信息
	CommentList []CommentInfo `json:"commentList"` //页面显示的评论信息
}

type IssueInfo struct {
	Issueid int64 `json:"issueid"`
	IssueName string `json:"issueName"`
	UpdateTime string `json:"updateTime"`
	UserName string `json:"userName"`
	TagName string `json:"tagName"`
	MilestoneName string `json:"milestoneName"`
}
```
  


### 4. "按页展示里程碑"

1. 路由定义

- Url: /milestone
- Method: POST
- Request: `MilestoneReq`
- Response: `MilestoneResp`

2. 请求定义


```golang
type MilestoneReq struct {
	Page int64 `json:"page"` //页码
	PageSize int64 `json:"pageSize"` //当前页面显示的milestone数量
}
```


3. 返回定义


```golang
type MilestoneResp struct {
	List []MilestoneInfo `json:"list"`
}
```
  


### 5. "创建里程碑"

1. 路由定义

- Url: /milestone/new
- Method: POST
- Request: `MilestoneNewReq`
- Response: `MilestoneNewResp`

2. 请求定义


```golang
type MilestoneNewReq struct {
	Name string `json:"name"`
}
```


3. 返回定义


```golang
type MilestoneNewResp struct {
}
```
  


### 6. "操作单个里程碑/按里程碑显示提案"

1. 路由定义

- Url: /milestone/signal
- Method: POST
- Request: `MilestoneSignalReq`
- Response: `MilestoneSignalResp`

2. 请求定义


```golang
type MilestoneSignalReq struct {
	Milestoneid int64 `json:"milestoneid"`
	ReName string `json:"reName"`
	IfDelete bool `json:"ifDelete"`
}
```


3. 返回定义


```golang
type MilestoneSignalResp struct {
	Milestone MilestoneInfo `json:"milestone"`
	IssueList []IssueInfo `json:"issueList"`
}

type MilestoneInfo struct {
	Milestoneid int64 `json:"milestoneid"`
	Name string `json:"name"`
	UpdateTime string `json:"updateTime"`
}
```
  


### 7. "按页展示标签"

1. 路由定义

- Url: /tag
- Method: POST
- Request: `TagReq`
- Response: `TagResp`

2. 请求定义


```golang
type TagReq struct {
	Page int64 `json:"page"` //页码
	PageSize int64 `json:"pageSize"` //当前页面显示的tag数量
}
```


3. 返回定义


```golang
type TagResp struct {
	List []TagInfo `json:"list"`
}
```
  


### 8. "创建标签"

1. 路由定义

- Url: /tag/new
- Method: POST
- Request: `TagNewReq`
- Response: `TagNewResp`

2. 请求定义


```golang
type TagNewReq struct {
	Name string `json:"name"`
}
```


3. 返回定义


```golang
type TagNewResp struct {
}
```
  


### 9. "操作单个标签/按标签显示提案"

1. 路由定义

- Url: /tag/signal
- Method: POST
- Request: `TagSignalReq`
- Response: `TagSignalResp`

2. 请求定义


```golang
type TagSignalReq struct {
	Name string `json:"name"`
	ReName string `json:"reName"`
	IfDelete bool `json:"ifDelete"`
}
```


3. 返回定义


```golang
type TagSignalResp struct {
	Tag TagInfo `json:"tag"`
	IssueList []IssueInfo `json:"issueList"`
}

type TagInfo struct {
	Tagid int64 `json:"tagid"`
	Name string `json:"name"`
}
```
  


### 10. "根据关键词搜索issue　包括issue标题, tag, comment　"

1. 路由定义

- Url: /issue/search
- Method: POST
- Request: `IssueSearchReq`
- Response: `IssueSearchResp`

2. 请求定义


```golang
type IssueSearchReq struct {
	Keyword string `json:"keyword"` //搜索的关键词
}
```


3. 返回定义


```golang
type IssueSearchResp struct {
	List []IssueInfo `json:"list"` //返回的issue列表
}
```
  

