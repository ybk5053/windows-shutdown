FROM alpine:latest
ARG TARGETARCH
ENV ARCH=${TARGETARCH}

WORKDIR /app

RUN apk --no-cache add samba

COPY windows-shutdown-${TARGETARCH} ./

EXPOSE 3030

RUN chmod +x /app/windows-shutdown-${TARGETARCH}

CMD /app/windows-shutdown-${ARCH}
