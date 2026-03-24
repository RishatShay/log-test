Сервис состоит из четырех частей (описаны в docker-compose.yaml)
    1. `seq-ui` - Бэкенд интерфейса. Обрабатывает запросы от фронтенда, взаимодействует с базой и организует поиск.
    2. `seq-ui-fe` - Фронтенд. Сервер, который при запросе возвращает бразеру код фронтенда
    3. `seq-db --mode proxy` - Точка входа для логов. Принимает данные по протоколу HTTP (Bulk API), занимается маршрутизацией запросов и балансировкой.
    4. `seq-db --mode store` - Движок хранения. Занимается непосредственной записью данных на диск и индексацией.

Запустить можно `docker compose up -d`
Логи после запуска смотреть: `localhost:5173`



Логи можно писать напрямую, обращаясь к seq-ui по порту 9002 (в формате bulk). Например через curl
```
curl --request POST \
  --url http://localhost:9002/_bulk \
  --header 'Content-Type: application/json' \
  --data '{"index" : {"unused-key":""}}
{"k8s_pod": "app-backend-123", "k8s_namespace": "production", "k8s_container": "app-backend", "request": "POST", "request_uri": "/api/v1/orders", "message": "New order created successfully"}
{"index" : {"unused-key":""}}
{"k8s_pod": "app-frontend-456", "k8s_namespace": "production", "k8s_container": "app-frontend", "request": "GET", "request_uri": "/api/v1/products", "message": "Product list retrieved"}
{"index" : {"unused-key":""}}
{"k8s_pod": "payment-service-789", "k8s_namespace": "production", "k8s_container": "payment-service", "request": "POST", "request_uri": "/api/v1/payments", "message": "failed"}
'
```


Или еще я написал простой код в `generator.go`, который периодически кидает логи разных уровней в сервис логирования.
Его нужно отдельно запустить. Просто тупо `go run generator.go`