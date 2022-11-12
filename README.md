# task_list

用 gin(golang) + gorm(mysql) 實作的 task api

採用 clean archtecture 架構, 使用 wire 當依賴注入

使用方式
```
git clone <repo>
docker-compose up
```

# Swagger
啟動服務後, http://localhost:8000/swagger/index.html

# API

1. GET /tasks (list tasks)
```
response
{
   "result": [
      {
         "id": 1,
         "name": "name",
         "status": 0
      }
   ]
}
```


2. POST /task (create task) 
```
request 
{
   "name": "買晚餐"
}

response status code 201
{
   "result": {
      "name": "買晚餐",
      "status": 0,
      "id": 1
   }
}
```


3. PUT /task/<id> (update task) 
```
request
{
   "id": 1,
   "name": "買早餐",
   "status": 1
}

response status code 200 
{
   "result": {
      "name": "買早餐",
      "status": 1,
      "id": 1
   }
}
```

4. DELETE /task/<id> (delete task)
```
response status code 200
```

# 專案目錄
#### api - 放 api handler
#### api_test - api handler test
#### config - 初始化配置
#### docs - swagger file
#### driver - 初始化 db
#### internal - Service 實作 業務邏輯, Repository 實作與 db 交互
#### middlware - 中間件
#### migration - init database
#### models - 定義 db schema 及用來操作 db 交互的 struct
#### models/apireq - API_Request
#### models/apires - API_Response
#### pkg - helper 工具
#### route - api 路由

# Database Schema

![image](https://upload.cc/i1/2022/11/12/ODvV5k.png)


# 測試連結
```
<!-- 服務跑在 EC2 -->

curl http://ec2-54-179-180-232.ap-southeast-1.compute.amazonaws.com:8000/tasks

swagger 

http://ec2-54-179-180-232.ap-southeast-1.compute.amazonaws.com:8000/swagger/index.html
```
