FROM golang:1.25-alpine


WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o auth ./cmd

EXPOSE 8001

CMD ["./auth"]
