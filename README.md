
# Activity Logger In Go and Mysql

Run the following command to start your server with mysql
```
$ docker-compose up --build
```


#### CREATE Activity Log

REQUEST:

```
$ curl --request POST 'http://localhost:8081/api/v1/logs' \
--header 'Content-Type: application/json' \
--data-raw '{
    "source": "engineroom",
    "type": "RIDE_DURATION",
    "entity_id": "20BGIJH",
    "entity_type": "ride",
    "field": "ride_duration",
    "old_value": "20",
    "new_value": "20",
    "data": "",
    "description": "",
    "action_by": "imran@gmail.com"
}'
```
RESPONSE: 201
```
{
    "success": true
}
```

#### GET Activity Logs with Pagination

Request

```
$ curl --request GET 'http://localhost:8081/api/v1/logs?page=1&source=engineroom&type=RIDE_DURATION' \
--header 'Content-Type: application/json'
```

RESPONSE: 200
```
{
    "data": [
        {
            "id": 1,
            "source": "engineroom",
            "type": "RIDE_DURATION",
            "entity_id": "20BGIJH",
            "entity_type": "ride",
            "field": "ride_duration",
            "old_value": "20",
            "new_value": "20",
            "data": "",
            "description": "",
            "action_by": "imran@gmail.com",
            "created_at": "2020-02-21T06:32:41Z"
        }
    ],
    "pagination": {
        "per_page": 10,
        "current_page": 1,
        "total": 1
    }
}
```

#### UPDATE An Existing Activity Log

REQUEST:

```
$ curl --request PATCH 'http://localhost:8081/api/v1/logs/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "source": "engineroom",
    "type": "RIDE_DURATION",
    "entity_id": "20BGIJH",
    "entity_type": "ride",
    "field": "ride_duration",
    "old_value": "20",
    "new_value": "22",
    "data": "",
    "description": "",
    "action_by": "fahim@gmail.com"
}'
```
RESPONSE: 200
```
{
    "success": true
}
```

####  DELETE An Existing Activity Log

REQUEST:
```
curl --request DELETE 'http://localhost:8081/api/v1/logs/1' \
--header 'Content-Type: application/json'
```
RESPONSE: 200
```
{
    "success": true
}
```
