## API go with gin

up api  (http://localhost:8080/)
```bash
air

```
Project structure
```bash
gin-app/
├── main.go
├── db/
│   ├── db.go
│   └── queries.go
├── handlers/
│   └── product_handlers.go  // Nuevo archivo para los handlers
└── models/
    └── product.go

```

build docker-compose (contain image api and db)

```bash
 docker-compose up -d
 ```

 Data bse in PgAdmin 
 ```bash
 http://localhost:5050

 us=admin@example.com
 pas=admin

 ```