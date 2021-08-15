FROM golang:alpine

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

RUN go get github.com/githubnemo/CompileDaemon

# add wait-for.sh script file
COPY wait-for.sh .

EXPOSE 8000

# **** following Entrypoint moved to docker-compose after adding wait-for script ****
# ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
