# Run dockerized
Set up **.env** file based on **.env.example** <br>
to start all containers run:
```sh
docker-compose up
```
*Note:* for best backend development experience (to enable hot golang rebuild), project folder should be mounted in ubuntu filesystem `~/project`, not windows `/mnt/c/user/project`



# Generate swagger docs
(having go-lang setup on local machine)
In ./backend/ dir run
```sh
swag init
```
Swagger docs are available at:
http://localhost:8000/swagger/index.html#/

# Frontend || react app

In order to access local (dev) api from local react add
```sh
"proxy": "http://api:8000",
```
 to /frontend/package.json 

`note:` *api* is the name of docker-compose container