FROM golang:1.13-alpine AS build
RUN apk add build-base
WORKDIR /go/src/github.com/org/repo
COPY . .

RUN go build -o server .

FROM alpine:3.12
EXPOSE 8080
COPY --from=build /go/src/github.com/org/repo/server /server
CMD ["/server"]