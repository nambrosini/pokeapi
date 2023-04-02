FROM golang:latest as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app .

FROM alpine:latest
COPY --from=build /app/app /app
CMD ["/app"]
