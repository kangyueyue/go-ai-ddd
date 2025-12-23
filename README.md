# go-ai-ddd
采用DDD（领域驱动设计）和AI技术的Go语言项目

`Application->Domain<-Infrastructure`

## 接口
### ping 服务活性测试
`curl "http://localhost:9091/api/v1/ping"`


### register 注册接口测试

`curl -X POST "http://localhost:9091/api/v1/user/register" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "123456",
    "captcha": "abc123"
  }'
`

### login 登入接口测试
`curl -X POST "http://localhost:9091/api/v1/user/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "user@example.com",
    "password": "123456"
  }'
`

### captcha 发送验证码接口
`curl -X POST "http://localhost:9091/api/v1/user/captcha" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com"
  }'
`