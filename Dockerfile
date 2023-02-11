FROM golang:1.18-alpine AS build
RUN apk --no-cache add build-base

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN mkdir bin
RUN go build -o /app/bin/rnt cmd/main.go


FROM alpine:3
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app/bin/rnt /app/rnt
COPY ./artifacts /app/artifacts

RUN chmod +x /app/rnt

CMD ["/app/rnt"]