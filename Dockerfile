FROM golang:1.15
ARG TARGETPLATFORM

COPY ./ui/public /ui
COPY ./build/${TARGETPLATFORM}/filer /filer
COPY ./entrypoint.sh /

WORKDIR /
CMD ["/entrypoint.sh"]
