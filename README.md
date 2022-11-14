# Task List API
[![CI](https://github.com/jtjin/task_list/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/jtjin/task_list/actions/workflows/ci.yml)

A Web service Task List (API) project designed by Go Language and using MySQL as a database for storing and managing tasks.

## How to run

You can run this app using 2 methods:

1. **Manual** (MySQL required)
    1. Clone this repo
    2. Run `go run main.go`
 
2. **Docker compose**
    1. Clone this repo.
    2. Run `docker-compose up -d` and wait for app to run on `localhost:8000`

### Demo
```
curl http://ec2-54-179-180-232.ap-southeast-1.compute.amazonaws.com:8000/tasks

http://ec2-54-179-180-232.ap-southeast-1.compute.amazonaws.com:8000/swagger/index.html
```

# Technologies
### Backend
- Languange: GO 1.19.x
- Web Framework: [gin 1.8.1](https://github.com/gin-gonic/gin)
- Object Relational Mapping: [gorm 1.24.1](https://github.com/go-gorm/gorm)
- Dependency Injection: [wire 0.5.0](https://github.com/google/wire)
- API Doc: [swag 1.8.7](https://github.com/swaggo/swag)
- Mocking Framework: [gomock 1.4.4](https://github.com/golang/mock)

### Cloud Service
- AWS EC2

### Database
- MySQL

### Tools
- Version Control: Git, Gitub
- CI / CD: Github Actions

### Others
- Design Pattern: Clean Architecture

# API Routes

- **Get all tasks**
   - End Point: `/tasks`
   - Method: GET
   - Request Example: `http://[HOST_NAME]/tasks`
   - Success Response: 200
      | Field | Type | Description |
      | :---: | :---: | :--- |
      | result | Array | Array of `Task Object` |
   - Success Response Example:
      ```json
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

- **Create a new task**
   - End Point: `/task`
   - Method: POST
   - Request Example: `http://[HOST_NAME]/task`
   - Request Headers:
      | Field | Type | Description |
      | :---: | :---: | :---: |
      | Content-Type | String | Only accept `application/json`. |    
   - Request Body:
      | Field | Type | Description |
      | :---: | :---: | :---: |
      | name (Required) | String | name of task |
      | status (Optional) | Number | status of task, only accept 0(incomplete) or 1(complete) |
   - Request Body Example:
      ```json
      {
         "name": "買晚餐"
      }
      ```   
   - Success Response: 201
      | Field | Type | Description |
      | :---: | :---: | :--- |
      | result | `Task Object` | Task information |
   - Success Response Example:
      ```json
      {
         "result": {
            "name": "買晚餐",
            "status": 0,
            "id": 1
         }
      }
      ```

- **Update/Edit a task**
   - End Point: `/task/<id>`
   - Method: PUT
   - Path Parameters
      | Field | Type | Description |
      | :---: | :---: | :--- |
      | id (Required) | Number | Task id |   
   - Request Example: `http://[HOST_NAME]/task/1`
   - Request Headers:
      | Field | Type | Description |
      | :---: | :---: | :---: |
      | Content-Type | String | Only accept `application/json`. |    
   - Request Body:
      | Field | Type | Description |
      | :---: | :---: | :---: |
      | name (Required) | String | name of task (Required) |
      | status (Required) | Number | status of task, only accept 0(incomplete) or 1(complete) |
   - Request Body Example:
      ```json
      {
         "name": "買早餐",
         "status": 1
      }
      ```   
   - Success Response: 200
      | Field | Type | Description |
      | :---: | :---: | :--- |
      | result | `Task Object` | Task information |
   - Success Response Example:
      ```json
      {
         "result": {
            "name": "買早餐",
            "status": 1,
            "id": 1
         }
      }
      ```

- **Delete a task**
   - End Point: `/task/<id>`
   - Method: Delete
   - Path Parameters
      | Field | Type | Description |
      | :---: | :---: | :--- |
      | id (Required) | Number | Task id |   
   - Request Example: `http://[HOST_NAME]/task/1`  
   - Success Response: 200

# Project Structure
```
├── .github
│   └──workflows          // github action 腳本
├── api                   // api handler
├── api_test              // api 整合測試
├── cmd
│   └──seeder             // 新增測試假資料的程式進入點
├── config                // 配置文件
├── docs                  // swagger file(swag init 自動生成)
├── driver                // 初始化資料庫
├── internal              
│   └── task           
│        ├── repository   // 實作與資料庫交互邏輯與單元測試
│        └── service      // 實作業務邏輯與單元測試
├── middleware            // 中介層
├── migration             // 資料庫初始化腳本(供 docker 使用)
├── mock                  // mock file(mockgen 自動生成)
├── models                // 定義 db schema 及用來操作 db 交互的 struct
│   ├── apireq            // 定義 API request struct
│   └── apires            // 定義 API response struct
├── pkg          
│   ├── errors            // 自定義錯誤 struct
│   ├── helper            // 輔助工具(ex: 生成 random 字串)
│   └── seeds             // 測試用的假資料
└── route                 // api 路由
```


# Database Schema

![image](https://upload.cc/i1/2022/11/12/ODvV5k.png)
