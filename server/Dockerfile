FROM golang:1.23.4-alpine

RUN apk add --no-cache git tzdata

ENV TZ=Asia/Jakarta

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd/main.go

CMD ["go", "run", "./cmd/main.go"]