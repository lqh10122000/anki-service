FROM golang:1.24.4-alpine

WORKDIR /app

# Install git (required by some go modules)
RUN apk update && apk add --no-cache git

COPY go.mod go.sum ./
ENV GOPROXY=direct
RUN go mod download

COPY . .

RUN go build -o main .

CMD ["./main"]
