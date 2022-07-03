# go-mock-api

How to setup

1) resources/application.yml

1.1 RDS 
setup:
  databaseType: "rds" or "dynamo" or "menkv"
  responseTime: 8900
  responseStatusCode: 200
  isRandomTime: 60
  count: 0

1.2 Dynamo

awsenv:
  aws_region: ""
  aws_access_id: ""
  aws_access_secret: ""

1.3 MenKv

GET http://localhost:8900/balance/list
GET http://localhost:8900/balance/id=6
GET http://localhost:8900/info
POST http://localhost:8900/cpu
    {
        "count":2000
    }

GET http://localhost:8900/balance/list_by_id/id=500&sk=order-5
POST http://localhost:8900/balance/save
    {
        "balance_id": "500",
        "account": "order-5",
        "amount": 500,
        "date_balance": "2020-01-01T00:00:00Z",
        "description": "sku-500"
    }

POST http://localhost:8900/setup
    {
        "response_time" : 1,
        "response_status_code" : 200,
        "is_random_time": true,
        "count":0
    }