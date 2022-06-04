#! /bin/bash
curl -X POST http://localhost:7712/api/users -d '{
    "name": "碩",
    "gender": "男",
    "nationID": "b07901052",
    "healthCardID": "123",
    "birthDay": "1999/12/11",
    "address": "新竹市民享街173巷13號",
    "phone": "0937988311",
    "vaccines": []
}
'