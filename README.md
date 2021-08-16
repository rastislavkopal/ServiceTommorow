# Run dockerized
Set up **.env** file based on **.env.example** <br>
to start all containers run:
```sh
docker-compose up
```



# Generate swagger docs
(having go-lang setup on local machine)
In ./backend/ dir run
```sh
swag init
```
Swagger docs are available at:
http://localhost:8000/swagger/index.html#/
