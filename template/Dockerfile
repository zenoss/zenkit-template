FROM golang:1.8.1-alpine

COPY . /go/src/{{PackageOwner}}/{{Name}}
WORKDIR /go/src/{{PackageOwner}}/{{Name}}/{{Name}}
RUN go build -o /bin/{{Name}}

FROM alpine
COPY --from=0 /bin/{{Name}} /bin/{{Name}}
COPY ./{{Name}}.yml /etc/{{Name}}.yml

EXPOSE {{Port}}

CMD ["/bin/{{Name}}", "server"]
