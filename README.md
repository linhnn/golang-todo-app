# Golang: Todo app

### Create table todo in Postgres
```
CREATE TABLE todo ( 
	id SERIAL PRIMARY key, 
	name VARCHAR(255), 
	date DATE 
)
```

### Run Go
```
go build 
./todo 
http://localhost:8080/list/ 
```
