# 生成api业务代码 ， 进入"服务/desc"目录下，执行下面命令
# goctl api go -api *.api -dir ../  --style=goZero


#protoc 生成 在rpc文件夹下
#goctl rpc protoc builder/user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.