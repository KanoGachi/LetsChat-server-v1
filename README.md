# LetsChat-server-v1
网页聊天小项目的服务端代码（Go）

## 数据库创建

服务端使用MySQL数据库，数据库中的表结构如下：

### user表

| 字段     | 类型    | 约束          | 备注                 |
| -------- | ------- | ------------- | -------------------- |
| uid      | bigint  | pk，auto incr | 用户id               |
| passport | bigint  | unique        | 登录口令             |
| name     | varchar |               | 名称                 |
| avasrc   | varchar |               | 头像路径             |
| status   | int     |               | 账号状态（保留字段） |

### friend表

| 字段  | 类型      | 约束                 | 备注             |
| ----- | --------- | -------------------- | ---------------- |
| uid   | bigint    | 联合主键（uid，fid） | id1              |
| fid   | bigint    | 联合主键（uid，fid） | id2              |
| ftime | timestamp |                      | 成为好友的时间点 |

### conv表

| 字段 | 类型      | 约束     | 备注         |
| ---- | --------- | -------- | ------------ |
| sid  | bigint    | not null | 发送者id     |
| rid  | bigint    | not null | 接收者id     |
| time | timestamp |          | 消息发送时间 |
| msg  | varchar   |          | 内容         |

### friendrequest表

| 字段      | 类型   | 约束                       | 备注                               |
| --------- | ------ | -------------------------- | ---------------------------------- |
| sessionid | bigint | pk，auto incr              | 请求编号                           |
| sid       | bigint | unique（uid，fid，status） | 发起者id                           |
| rid       | bigint | unique（uid，fid，status） | 请求者id                           |
| status    | int    | unique（uid，fid，status） | 请求状态（1接受，0未处理，-1拒绝） |


## 配置数据库地址

需要在`config/config.ini`中配置`mysqldsn="用户名:密码@tcp(IP地址:3306)/数据库名"`

## 启动服务器

```bash
go run main.go
```
