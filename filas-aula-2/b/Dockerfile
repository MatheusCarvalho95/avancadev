FROM golang:1.15

WORKDIR /go/src/serviceB

COPY . . 

RUN GOOS=linux go build

ENTRYPOINT [ "./serviceB" ]