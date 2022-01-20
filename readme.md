**Notes**<br>
If you haven't install go, you can install it <a href="https://go.dev/doc/install"> here </a><br>
If you are using Mac M1 Chip, you need to add ```platform: linux/x86_64``` on docker-compose.yaml before build the docker (put it on the same tab-line with image)

**Test 1**<br>
**Step By Step**
1. Run ```https://github.com/josechenn/golang-test-1.git``` on terminal
2. Run ```sudo cp .env.example .env```  to copy .env.example as .env
2. Run ```docker-compose build``` to create mysql server for running the code
4. Run ```go mod tidy``` to get missing and remove unused modules
5. Before migrate, we need to install goose, how to install goose can be found <a href="https://formulae.brew.sh/formula/goose"> here </a>
6. After install goose we can run ```goose mysql "user:password@tcp(127.0.0.1:3306)/db?parseTime=true" up``` to migrate current migrations
7. Run ```go run main.go seed``` to seed the database for testing purpose
8. Next we need to move to testing folder by using ```cd testing/```
9. Run ```go test -v``` to run all job, to run a specified job we need to use ```go test -v -run {test_name}```
10. Run ```go run main.go``` to access the API via postman/ARC/Etc (Api Route has been declared on main.go)<br>

<i>ps : Route Explanation can be read on main.go</i>


**Test 2**<br>
**Step By Step**
1. Run ```go run main.go```
2. Access ```http://localhost:1234/count_total``` to see the result (the number of apple can be changed on cake_service.go)
3. Re run ```go run main.go``` to update the code
