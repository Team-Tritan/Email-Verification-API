FROM golang

COPY * *

RUN go build main.go

RUN ./main