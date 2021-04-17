FROM golang:1.16.3-alpine3.12

ENV GO111MODULE=on
EXPOSE 8080

WORKDIR /myhttp
COPY go.mod .

RUN go mod download 
COPY . .
RUN go build 

ENTRYPOINT ["go"]
CMD [ "run", "main.go" ]
