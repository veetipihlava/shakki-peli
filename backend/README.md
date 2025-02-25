# Shakki-peli
Chess game course project for Aalto university  CS-E4003 - Special Assignment in Computer Science D

# testing
go test ./...


# Redis

The easiest way to run Redis is to pull it and run it as a container with Docker:  

`` docker pull redis ``  

`` docker run --name redis -p 6379:6379 -d redis ``

To test that Redis is working you can run commands inside it:  
`` docker exec -it redis redis-cli ``  
`` set mykey "chessgame" ``  
`` get mykey ``  
`` del mykey ``  