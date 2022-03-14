# fin

gin风格的微型web api框架，内置gorm

本框架仅供学习参考，不建议使用于生产环境

想学gin+gorm？传送门：https://github.com/feizhiwu/go-stage

### 运行demo
```
> git clone ...
> cd fin
> go run example/main.go
```

### RESTful API curl

```
添加用户：
curl --location --request POST 'http://localhost:8080/v1/user' \
--header 'action: add' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"tout",
    "password":"123"
}'

用户列表：
curl --location --request GET 'http://localhost:8080/v1/user?page=1&limit=10' \
--header 'action: list'

用户详情：
curl --location --request GET 'http://localhost:8080/v1/user?id=1' \
--header 'action: info'

修改用户:
curl --location --request PUT 'http://localhost:8080/v1/user' \
--header 'action: update' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":1,
    "password":"123456"
}'

删除用户：
curl --location --request DELETE 'http://localhost:8080/v1/user' \
--header 'action: delete' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":1
}'
```

注：一个router下的header action需与controller方法同名，联调时可以通过action快速定位接口

### 返参示例

```
{
    "status": 10000,
    "msg": "请求成功",
    "body": null
}
```

### 测试user表结构

```
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```
