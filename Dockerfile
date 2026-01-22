FROM golang:1.23-alpine AS dev
WORKDIR /app
RUN apk add --no-cache git
RUN go install github.com/air-verse/air@v1.60.0
COPY go.mod go.sum ./
RUN go mod download
COPY . .
CMD ["air", "-c", ".air.toml"]

FROM golang:1.23-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -buildvcs=false -o server ./cmd/server

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/server .
COPY --from=build /app/database ./database
COPY --from=build /app/static ./static

EXPOSE 7741
CMD ["./server"]