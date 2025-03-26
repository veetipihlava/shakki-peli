# Redis

The easiest way to run Redis is to pull it and run it as a container with Docker:  

`` docker pull redis `` 
`` docker run --name redis -p 6379:6379 -d redis ``

To test that Redis is working you can run commands inside it:  
`` docker exec -it redis redis-cli ``  
`` set mykey "chessgame" ``  
`` get mykey ``  
`` del mykey ``  




# Running with Docker compose

With Windows:
Docker desktop running & inside deployment directory:

Build and run services
`` docker compose up --build ``

Run services (and CTRL + C to stop services)
`` docker compose up ``

Stop and remove services
`` docker compose down ``