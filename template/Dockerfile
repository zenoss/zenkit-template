#
# Zenoss CONFIDENTIAL
# -------------------
#
# This software Copyright (c) Zenoss, Inc. 2025
# All Rights Reserved.
#
# NOTICE:  All information contained herein is, and remains the property of
# Zenoss Incorporated and its suppliers, if any.  The intellectual and
# technical concepts contained herein are owned and proprietary to Zenoss
# Incorporated and its suppliers and may be covered by U.S. and Foreign
# Patents, patents in process, and are protected by U.S. and foreign trade
# secret or copyright law. Dissemination of this information or reproduction of
# any this material herein is strictly forbidden unless prior written
# permission by an authorized officer is obtained from Zenoss Incorporated.
#
ARG ZENKIT_BUILD_VERSION=1.24.5
FROM gcr.io/zing-registry-188222/zenkit-build:${ZENKIT_BUILD_VERSION} AS builder

ARG GOAMD64=v3
ENV GOAMD64=$GOAMD64
COPY . /workspace/
WORKDIR /workspace/
RUN go build -mod vendor -o /bin/{{Name}}

FROM alpine:3
COPY --from=builder /bin/{{Name}} /bin/{{Name}}
RUN apk add --no-cache 'curl>=8' ; \
    addgroup -S -g 512 zing && adduser -S -G zing -u 512 zing

USER 512:512

EXPOSE 8080

CMD ["/bin/{{Name}}"]
