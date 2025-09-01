FROM golang:1.24

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o actions-for-go ./cmd

EXPOSE 8000

CMD ["./actions-for-go"]
