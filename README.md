# go-ai-ddd
采用DDD（领域驱动设计）和AI技术的Go语言项目

`Application->Domain<-Infrastructure`

## 接口
### register 注册测试

`curl -X POST "http://localhost:9091/api/v1/register" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "123456",
    "captcha": "abc123"
  }'
`
