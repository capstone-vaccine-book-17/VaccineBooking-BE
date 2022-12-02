FROM golang:latest as wvac

WORKDIR /app
COPY . .

RUN go build -tags netgo -o wvac-be .

FROM alpine:latest

COPY --from=wvac /app/wvac-be .
COPY --from=wvac /app/.env .

CMD ["/wvac-be"]