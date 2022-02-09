# go-mock-api

GET http://localhost:8900/balance/list
GET http://localhost:8900/balance/id=6
GET http://localhost:8900/info
POST http://localhost:8900/cpu
    {
        "count":2000
    }

GET http://localhost:8900/balance/list_by_id/id=6&sk=order-7
POST http://localhost:8900/balance/save
    {
        "balance_id": "6",
        "account": "order-7",
        "amount": 10,
        "date_balance": "2020-01-01T00:00:00Z",
        "description": "oiiii"
    }

POST http://localhost:8900/setup
    {
        "response_time" : 1,
        "response_status_code" : 200,
        "is_random_time": true,
        "count":0
    }