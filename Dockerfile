
FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY main.go ./

COPY ./controller ./controller
COPY ./domain ./domain
COPY ./data ./data

ENV PORT 3000
EXPOSE 3000

RUN go build -o /hotel-service

CMD [ "/hotel-service" ]
