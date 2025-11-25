# Prometheus POC

## Requests

### GET /api/todos

`curl http://localhost:8080/api/todos`

### POST /api/todos

```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"id": 3,"task":"third"}' \
  http://localhost:8080/api/todos
```
