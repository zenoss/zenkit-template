FROM golang:1.8.3-alpine
RUN apk --no-cache add git bash
RUN go get github.com/tmrts/boilr && \
	cd /go/src/github.com/tmrts/boilr && \
	git remote add iancmcc http://github.com/iancmcc/boilr && \
	git fetch iancmcc master:iancmcc-master && \
	git checkout iancmcc-master && \
	go install
COPY ./docker-entrypoint.sh /usr/local/docker-entrypoint.sh
ENTRYPOINT ["/usr/local/docker-entrypoint.sh"]
