**Test 1**
**Step By Step**
1. Run ```https://github.com/josechenn/golang-test-1.git``` on terminal
2. Run ```sudo cp .env.example .env```  to copy .env.example as .env
2. Run ```docker-compose build``` to create mysql server for running the code
4. Run ```go mod tidy``` to get missing and remove unused modules
5. Run ```goose mysql "root:root@tcp(127.0.0.1:3306)/db?parseTime=true" up``` to migrate current migrations
6. Run ```go run main.go seed``` to seed the database for testing purpose
7. Next we need to move to testing folder by using ```cd testing/```
8. Run ```go test -v``` to run all job, to run a specified job we need to use ```go test -v -run {test_name}```
9. Run ```go run main.go``` to access the API via postman/ARC/Etc (Api Route has been declared on main.go)

***Route Explanation can be read on main.go***


**Test 2**
1. Run ```go run main.go```
2. Access ```http://localhost:1234/count_total``` to see the result (the number of apple can be changed on cake_service.go)
3. Re run ```go run main.go``` to update the code