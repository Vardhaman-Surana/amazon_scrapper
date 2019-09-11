FROM golang:1.12 AS builder
WORKDIR /go/src/github.com/vds/amazon_scrapper
ENV GO111MODULE=on
COPY ./go.mod ./go.sum ./
RUN go mod download
RUN go mod verify
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/Scrapper ./cmd/ScrapperService/main.go

FROM alpine:latest AS scrapper
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash
COPY --from=builder /go/src/github.com/vds/amazon_scrapper/bin/Scrapper .
CMD ["./Scrapper"]