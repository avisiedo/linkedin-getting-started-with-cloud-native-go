FROM docker.io/golang:1.15.1-alpine AS binary

ENV SOURCES /go/src/github.com/avisiedo/go-cloud-native/
COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build -mod vendor -o microservice



FROM docker.io/golang:1.15.1-alpine

COPY --from=binary /go/src/github.com/avisiedo/go-cloud-native/microservice /microservice

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT [ "/microservice" ]
