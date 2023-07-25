# GOLANG TECHNICAL TEST

## TEST CASE 3
a. Munculkan data country mana aja yang spend nya terbanyak (Query)<br/>
b. Munculkan data jumlah tipe kartu kredit terbanyak (Query)<br/>
c.  Buatlah GET API untuk medapatkan data dari soal a dengan response menampilkan :
* Id 
* Country 
* Credit_card_type 
* Credit_card 
* First_name 
* Last_name 

d. Buatlah POST API untuk mengirim data dengan body seperti dibawah ini : 
```shell
{ 
    "country”: " ", 
    "credit_card_type”: " ", 
    "credit_card”: " ", 
    "first_name”: " ", 
    "last_name”: " " 
} 
```

## Endpoint
3.a. http://localhost:9090/api/v1/get-country-spend<br/>
3.b. http://localhost:9090/api/v1/get-total-cc-by-type<br/>
3.c. http://localhost:9090/api/v1/:country/get-data<br/>
3.d. http://localhost:9090/api/v1/create-users<br/>


## Requirements
- Golang Echo
- PostgreSql

## Installation
### Local
```shell
//change directory to project
cd golangtest
//downloading go module dependency
go mod tidy
//running project 
go run main.go
```

## Migration Files
- Using `AutoMigrate` from [gorm package](https://gorm.io/docs/migration.html)
## JSON Seeder Files
Using Auto Seeder 
- resources/datasources/users.json
- resources/datasources/belanja.json
