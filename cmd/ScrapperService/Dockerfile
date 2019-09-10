FROM golang:1.12-alpine
LABEL maintainer="Vardhaman Surana"
WORKDIR ./amazon_scrapper
COPY ./go.mod  ./go.sum ./
RUN apk add  git
RUN go mod download
RUN go mod verify
COPY ./ ./
RUN go build -o ./bin/Scrapper ./cmd/ScrapperService
CMD ["./bin/Scrapper"]

