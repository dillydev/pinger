FROM alpine:3.5

RUN apk --no-cache add ca-certificates

ARG EXEC_NAME

ADD $EXEC_NAME /bin/server