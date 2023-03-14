FROM golang:latest as builder

WORKDIR /app

COPY ola-mundo.go /app 

RUN go build ola-mundo.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/ola-mundo .

ENTRYPOINT [ "./ola-mundo" ]