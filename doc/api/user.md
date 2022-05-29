
### 1. "用户注册"

1. 路由定义

- Url: /user/register
- Method: POST
- Request: `RegisterReq`
- Response: `RegisterResp`

2. 请求定义


```golang
type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
```


3. 返回定义


```golang
type RegisterResp struct {
	AccessToken string `json:"accessToken"` //令牌
	AccessExpire int64 `json:"accessExpire"` //令牌过期时间
	RefreshAfter int64 `json:"refreshAfter"`
}
```
  


### 2. "用户登录"

1. 路由定义

- Url: /user/login
- Method: POST
- Request: `LoginReq`
- Response: `LoginResp`

2. 请求定义


```golang
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
```


3. 返回定义


```golang
type LoginResp struct {
	AccessToken string `json:"accessToken"` //令牌
	AccessExpire int64 `json:"accessExpire"` //令牌过期时间
	RefreshAfter int64 `json:"refreshAfter"`
}
```
  


### 3. "获取用户信息"

1. 路由定义

- Url: /user/info
- Method: POST
- Request: `GetUserInfoReq`
- Response: `GetUserInfoResp`

2. 请求定义


```golang
type GetUserInfoReq struct {
	Userid int64 `json:"userid"`
	Mobile string `json:"mobile"`
	Username string `json:"username"`
}
```


3. 返回定义


```golang
type GetUserInfoResp struct {
	Userid int64 `json:"userid"`
	Mobile string `json:"mobile"`
	Username string `json:"username"`
}
```
  

