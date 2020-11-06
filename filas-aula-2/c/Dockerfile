FROM golang:1.15

WORKDIR /go/src/serviceC

COPY . . 

RUN GOOS=linux go build

EXPOSE 9092

ENTRYPOINT [ "./serviceC" ]