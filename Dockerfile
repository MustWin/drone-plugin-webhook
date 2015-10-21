FROM gliderlabs/alpine:3.1
RUN apk-install ca-certificate
ADD drone-plugin-webhook /bin/
ENTRYPOINT ["/bin/drone-plugin-webhook"]
