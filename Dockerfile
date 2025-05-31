FROM golang:1.24 as builder
WORKDIR /usr/src/app
COPY . .
RUN go mod download
RUN go build ./cmd/dark-summoner-be

FROM golang:1.24
COPY --from=builder /usr/src/app/dark-summoner-be /usr/local/bin

CMD ["dark-summoner-be"]