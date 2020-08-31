FROM golang:1.14.3-alpine as build

COPY . /app
WORKDIR /app
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.io,direct
RUN go build -o server main.go
RUN chmod +x ./server

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/server /app
COPY ./config.yaml /app
EXPOSE 8080
CMD ["./server"]