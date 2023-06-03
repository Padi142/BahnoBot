FROM golang:1.20.3-alpine3.17 as builder

ADD . /go/src/bahno_bot

WORKDIR /go/src/bahno_bot

RUN go mod tidy

RUN go build -o /build/bahno_bot

RUN go install github.com/swaggo/swag/cmd/swag

RUN swag init

FROM alpine:3.17.3 as final

WORKDIR /bahno_bot

COPY --from=builder /build/bahno_bot /bahno_bot

COPY ./.env /bahno_bot

EXPOSE 8081

CMD [ "/bahno_bot/bahno_bot" ]




