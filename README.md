# go-url-shortener

---
# Software Architecture
![image](./doc/img/architecture.png)

---
# API list

### 1. /api/ping 
Request Method: GET

Response:
```json
{
    "status": "up and running"
}
```

### 2. /api/shorten
Request Method: POST

Request Body:
```
{
  "url": "https://www.youtube.com/watch?v=TLB5MY9BBa4&ab_channel=CoderDave"
}
```
Response:
```json
{
    "results": {
        "id": 13,
        "actual_url": "https://www.youtube.com/watch?v=TLB5MY9BBa4&ab_channel=CoderDave",
        "short_url": "77cf66ac",
        "total_hit": 0,
        "created_at": "2023-07-23T14:37:53.241159Z"
    }
}
```

### 3. /api/decode/{shortUrl}
Request Method: GET
Response:
```json
{
    "results": {
        "id": 13,
        "actual_url": "https://www.youtube.com/watch?v=TLB5MY9BBa4&ab_channel=CoderDave",
        "short_url": "77cf66ac",
        "total_hit": 6,
        "created_at": "2023-07-23T14:37:53.241159Z"
    }
}
```
