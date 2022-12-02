FROM golang:latest as wvac

WORKDIR /app
COPY . .

RUN go build -tags netgo -o wvac-be .

FROM alpine:latest

RUN apk update && apk add --no-cache git

COPY --from=wvac /app/wvac-be .

CMD ["/wvac-be"]