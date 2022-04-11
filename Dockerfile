FROM golang

WORKDIR /go/src/

COPY . .

RUN go get .

RUN go build -o main

EXPOSE 8080

CMD ["./main"]