FROM golang:1.21-alpine3.18  as builder

RUN apk add git

LABEL maintainer="7fantasy7@gmail.com"

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# GO Repo base repo
FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

RUN mkdir /app

WORKDIR /app/

COPY --from=builder /app/main .

# Expose port 8000
EXPOSE 8000

CMD ["./main"]