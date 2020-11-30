# Mansion Assessment
GoLang Micro service for zip code location filtering

Example for testing
```
curl -X POST -d "{\"postcode\": \"AL9 5JP\", \"radius\": 30}" http://localhost:8080
```

```
POST http://localhost:8080
Content-Type: application/json

{
    "postcode": "AL9 5JP",
    "radius": 30
}

###
```