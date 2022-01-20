**Step By Step**
1. Run ```https://github.com/josechenn/golang-test-1.git``` on terminal
2. Run ```docker-compose build``` to create mysql server for running the code
4. Run ```go mod tidy``` to get missing and remove unused modules
5. Run ```goose mysql "root:root@tcp(127.0.0.1:3306)/db?parseTime=true" up``` to migrate current migrations
6. Run ```go run main.go seed``` to seed the database for testing purpose
7. Next we need to move to testing folder by using ```cd testing/```
8. Run ```go run test -v``` to run all job, to run a specified job we need to use ```go run -v -run {task_name}```

***Route Explanation can be read on main.go***