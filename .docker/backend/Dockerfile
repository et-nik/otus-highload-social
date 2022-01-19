FROM golang:alpine as builder

ENV APPDIR $GOPATH/src/github.com/et-nik/otus-highload-social

RUN mkdir -p ${APPDIR}
WORKDIR ${APPDIR}

COPY ./go.mod .
COPY ./go.sum .

RUN go version \
  && go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app ./cmd/social

FROM alpine:3.7

COPY --from=builder /app /app

ENV PORT=8000 \
    DATABASE=root:password@(localhost:3306)/social

EXPOSE 8000

CMD ["./app"]
