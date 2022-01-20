**FIRST**
1. Run ```docker-compose build```
2. Run ```goose mysql "user:password@tcp(127.0.0.1:3306)/db?parseTime=true" up``` to migrate current migrations
3. 