# Dockerfile.production

FROM registry.semaphoreci.com/golang:1.18 as builder

ENV APP_HOME /go/src/server

WORKDIR "$APP_HOME"
COPY src/ .

RUN go mod download
RUN go mod verify
RUN go build -o energy server/server.go

FROM registry.semaphoreci.com/golang:1.18

ENV APP_HOME /go/src/
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY src/website/ /go/website/
COPY --from=builder "$APP_HOME"server/energy $APP_HOME

EXPOSE 80
CMD ["./energy"]