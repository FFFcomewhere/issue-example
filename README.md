## 功能实现
本项目是一个模仿github的issue，的任务管理程序，
1. 支持用户创建，登录账号
2. 操作issue,milestone,tag,comment包括创建，修改和删除
3. 根据页码，tag,milestone列出issue
4. 根据用户输入的关键词，搜索issue标题，tag，comment找到相关的issue


## 技术栈
- go-zero 微服务框架
- docker 容器化
- docker-compose 环境部署＋服务部署
- mysql 持久化数据库
- redis　缓存
- elasticsearch　文本搜索

## 目录介绍

- app:业务代码的api和rpc
- common:通用组件error,tool等
- data:项目依赖的中间件（mysql,redis,es），该目录在git忽略文件内，无需提交
- doc:
  - api:各个api接口的请求路径，请求体和相应体，以及注意事项。
- deploy:部署环境的依赖文件和脚本
  - esves的ik分词器，需要手动移动到data/plugins文件夹内
  - goctl:项目的goctl自定义模块，如果需要重新生成代码，请复制到本地.goctl文件夹内，具体用法参考go-zero文档
  - mysql2es:go-mysql-elasticsearch组件的配置文件，功能是将mysql内数据同步至es
  - script:
    - gencode:代码生成脚本
    - mysql:生成model的脚本
  - sql:sql的ddl文件，在数据库初始化完成后，需要手动运行
- .env: docker-compose的配置文件
- docker-compose.yml: 构建和部署服务的配置文件
- docker-compose-env.yml: 构建和部署环境的配置文件


### 业务代码目录
这里以用户服务为例
- usercenter:用户中心服务
  - api:对外api接口
    - builder:api文件
    - etc:服务中间件配置文件
    - internal:核心代码
      - config:配置结构体
      - handler:路由接口
      - logic:业务逻辑
      - svc:中间件注入
      - types:请求，响应结构体
    - Dockerfile:可自动化生成的dockerfile
    - usercenter.go:主程序,main
  - model:数据库modle结构体和CURD函数
  - rpc:　对内rpc服务接口,整体结构和api目录类似，不再赘述




## 环境部署

### 安装docker和docker-compose, 部署组件

运行docker-compose-env.yml文件，依赖环境可能不能一次运行起来，自动生成的data文件夹可能是只读的，elasticsearch无法读取，所以需要把data目录下的子文件夹设置为可读可写。

运行sudo chmod 777 -R data/指令，修改完权限后重写启动es,然后再将deploy/es/ik文件夹，复制到/data/es/plugins内 ,重启es。


为了后面能把mysql的数据传到es中，mysql需要打开binlog ,在命令行内输入下列命令，并重启mysql。

docker exec mysql bash -c "echo 'log-bin=/var/lib/mysql/mysql-bin' >> /etc/mysql/mysql.conf.d/mysqld.cnf"

docker exec mysql bash -c "echo 'binlog-format=Row' >> /etc/mysql/mysql.conf.d/mysqld.cnf"

docker exec mysql bash -c "echo 'server-id=123454' >> /etc/mysql/mysql.conf.d/mysqld.cnf"



### 初始化数据库
进入/deploy/sql内，运行sql，先运行usercenter.sql,再运行issueceter.sql。
注意用户的密码加密过，明文密码为password。
等表格创建好后，重启go-mysql-elasticesearch容器。


### 部署并运行服务

运行rpc和api，所有服务都由dockerfile构建，并由docker-compose统一部署，
直接运行docker-compose.yml文件，构建需要花一点时间，请耐心等待。

三者的依赖关系

usercenter-rpc -> usercenter-api


usercenter-rpc -> issuecenter-api

### 测试
部署完环境之后，可以进行测试，测试文件在app下各个服务的目录下，
建议先阅读doc内的api文件,对比该文件进行测试。

