FROM golang:1.24-alpine AS build

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o /app/server ./server

FROM alpine:latest

RUN apk add --no-cache sqlite-libs

WORKDIR /root/

COPY --from=build /app/server .

EXPOSE 8080

CMD ["./server"]
