FROM alpine:latest
ARG TARGETARCH

WORKDIR /app

RUN apk add samba

COPY windows-shutdown-${TARGETARCH} ./

EXPOSE 3030

RUN chmod +x /app/windows-shutdown-${TARGETARCH}

CMD /app/windows-shutdown-${TARGETARCH}
