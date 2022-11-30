# aiops

## 一、初始化
### 1.创建目录
```shell
mkdir aiops && cd aiops
mkdir app common deploy doc data
go mod init aiops
```

### 2.修改response模板
参考链接：https://go-zero.dev/cn/docs/advance/template

### 3.创建user服务
```shell
mkdir -p app/user/api 
mkdir -p app/user/rpc
touch app/user/api/user.api
touch app/user/rpc/user.proto
```
```shell
cd app/user/rpc
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
```
生成usermodel
```shell
cd app/user/model
goctl model mysql ddl -src ../../../deploy/sql/user.sql -dir . -c
```

### 4.创建duty服务
> 值班管理服务
```shell
mkdir -p app/duty/api 
mkdir -p app/duty/rpc
touch app/duty/api/duty.api
touch app/duty/rpc/duty.proto

```
```shell
cd app/duty/api
goctl api go -api duty.api -dir .
```