FROM golang:latest

#RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go mod tidy

#RUN make build

RUN cd ./server/grpc && go build

EXPOSE 9902

ENTRYPOINT ["./main"]