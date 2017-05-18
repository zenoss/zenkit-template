FROM golang:1.8.1-alpine

COPY . /go/src/{{Package}}
WORKDIR /go/src/{{Package}}/{{Name}}
RUN go build -o /bin/{{Name}}

FROM alpine
COPY --from=0 /bin/{{Name}} /bin/{{Name}}
COPY ./{{Name}}.yml /etc/{{Name}}.yml
CMD ["/bin/{{Name}}", "server"]
