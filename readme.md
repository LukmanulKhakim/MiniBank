## Mini Bank API

Repository ini berisi tentang Mini Bank API yaitu API yang menunjukkan simulasi sistem perbankan, yang terdiri atas credit dan debit pada sebuah akun bank.
Dalam API ini terdapat 2 Fitur 
1. Account
2. Transaksi

### Cara Run Local

1. Clone Repo
```
git clone https://github.com/LukmanulKhakim/MiniBank

cd minibank
```
2. Create .env file
```
DB_USER = root
DB_PWD = password mysql
DB_HOST = localhost
DB_PORT = 3306
DB_NAME = minibank -- buat database

```
3. run 
```
go run server.go
```

### Routing
```
POST /account
GET  /account/list?details=true
GET  /account/:account_no
GET  /account/list?details=false

POST /transaction
GET  /transaction/list
GET  /transaction
POST /transaction/search
```
