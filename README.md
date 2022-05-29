# 环境部署

1. 安装docker和docker-compose, 运行docker-compose-env.yml文件

注意：依赖环境可能不能一次运行起来，自动生成的data文件夹可能是只读的，elasticsearch无法读取，所以需要把data目录下的子文件夹设置为可读可写
例如:　 sudo chmod 777 -R /data
修改完权限后重写启动es
然后再将deploy/es/ik文件夹　复制到/data/es/plugins内
重启es


为了后面能把mysql的数据传到es中，mysql需要打开binlog
docker exec mysql bash -c "echo 'log-bin=/var/lib/mysql/mysql-bin' >> /etc/mysql/mysql.conf.d/mysqld.cnf"
docker exec mysql bash -c "echo 'binlog-format=Row' >> /etc/mysql/mysql.conf.d/mysqld.cnf"
docker exec mysql bash -c "echo 'server-id=123454' >> /etc/mysql/mysql.conf.d/mysqld.cnf"
重启mysql


2. 运行sql 　在/deploy/sql内，依次运行两个文件
注意　用户的密码过　明文密码为password
等表格创建好后，重启go-mysql-elasticesearch容器

3. 运行rpc 和　api，　所有服务都由dockerfile构建，并由docker-compose统一部署
直接运行docker-compose.yml文件，构建需要花一点时间，请耐心等待
三者的依赖关系
usercenter-rpc -> usercenter-api
usercenter-rpc -> issuecenter-api