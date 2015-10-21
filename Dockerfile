FROM gliderlabs/alpine:3.1
RUN apk-install ca-certificate
ADD drone-webhook /bin/
ENTRYPOINT ["/bin/drone-webhook"]
