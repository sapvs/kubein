FROM golang:alpine as build
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . . 
RUN go build -o server cmd/main.go

FROM alpine 
WORKDIR /app
COPY --from=build /app/server .
ENTRYPOINT [ "./server" ]