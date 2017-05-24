# demo
a demo of golang


简单实现了用户的注册，登录，修改密码，删除用户的接口
简单记录了方法调用日志

demo中为了方便，全都使用了GET method，没有使用POST,DELETE,PUT 等

go run main.go
或者  go build


4个接口地址
http://localhost:8888/users/register?username=test&password=test
http://localhost:8888/users/login?username=test&password=test
http://localhost:8888/users/changepassword?username=test&password=test&newpass=test1
http://localhost:8888/users/delete?username=test&password=test

