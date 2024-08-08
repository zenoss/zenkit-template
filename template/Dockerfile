ARG ZENKIT_BUILD_VERSION=1.18.2

FROM zenoss/zenkit-build:${ZENKIT_BUILD_VERSION}

COPY . /workspace/
WORKDIR /workspace/
RUN go build -v -mod vendor -o /bin/{{Name}}

FROM alpine
RUN apk add --no-cache curl

COPY --from=0 /bin/{{Name}} /bin/{{Name}}
RUN addgroup -S zing && adduser -S -G zing -u 512 512
USER 512

EXPOSE {{Port}}

CMD ["/bin/{{Name}}"]
