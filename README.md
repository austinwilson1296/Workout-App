# Fitted API

## Project Description

Api to generate workouts for all levels of fitness experience.

## Usage

### Create User

``` #bash

curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "username": "testuser",
    "password": "password123",
    "experience_level_id": 2
  }' 

```

### Login User

``` #bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123",
  }' 
```

### Generate WarmUp

>Level defines what exercises will be generated.
>
>Exercises are selected randomly and returned in the following order:

1. Core Hips Legs
2. Core Hips Legs
3. Core Spinal
4. Core Spinal
5. Thoracic Spine Mobility
6. Thoracic Spine Mobility
7. Scapulo Thoracic
8. Scapulo Thoracic
9. Shoulders Scapula
10. Shoulders Scapula

``` #bash
curl http://localhost:8080/api/warmup?level=1
```

### Generate Main Exercise

>Generates exercise based on provided level and body part

``` #bash
curl http://localhost:8080/api/exercise?level=1&name=chest
```
