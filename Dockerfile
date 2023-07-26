# FROM golang:1.20.3-alpine3.17 as builder
FROM golang:1.20-bullseye as builder

ADD . /bahno_bot
WORKDIR /bahno_bot

RUN apt -y update && apt -y install python3 python3-venv python3-pip
RUN python3 -m venv /opt/venv
ENV PATH="/opt/venv/bin:$PATH"
COPY ./requirements.txt .
RUN pip install -r requirements.txt

RUN go mod tidy
RUN go build -o /build/bahno_bot
RUN go get github.com/swaggo/swag/cmd/swag@v1.16.1
RUN go get github.com/swaggo/swag/gen@v1.16.1
RUN go install github.com/swaggo/swag/cmd/swag

RUN swag init

FROM debian:bullseye-slim as final
WORKDIR /bahno_bot

COPY --from=builder /build/bahno_bot /bahno_bot
COPY --from=builder /opt/venv /opt/venv

COPY ./.env.example ./.env* /bahno_bot/
COPY ./docs/ /bahno_bot/docs
COPY ./charts.py ./charts.py

ENV PATH="/opt/venv/bin:$PATH"

RUN apt update && apt -y install ca-certificates python3 && update-ca-certificates 

EXPOSE 8081

CMD [ "/bahno_bot/bahno_bot" ]