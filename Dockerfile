FROM golang:1.17.6-alpine3.15

WORKDIR /app

RUN apk update --no-cache \
  && apk add --no-cache \
  git \
  gcc \
  musl-dev

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o app src/main.go

RUN GO111MODULE=off go get github.com/oxequa/realize
RUN GO111MODULE=off go get github.com/smartystreets/goconvey
RUN GO111MODULE=off go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate

EXPOSE 8000

# CMD [ "./app" ]