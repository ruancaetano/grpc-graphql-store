FROM golang:1.19.0-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o ./bin/server ./cmd/server.go


FROM golang:1.19.0-alpine as runner

WORKDIR /app

COPY --from=builder ./app/bin/server .
COPY --from=builder ./app/.env .

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait ./wait
RUN chmod +x ./wait

EXPOSE 8001

CMD ./server