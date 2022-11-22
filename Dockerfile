FROM golang:1.18.2-alpine as build

COPY build/app/go.mod src/go.mod
COPY build/app/cmd src/cmd/
COPY build/app/models src/models/
COPY build/app/restapi src/restapi/

RUN cd src/ && go mod tidy

RUN cd src && \
    export CGO_LDFLAGS="-static -w -s" && \
    go build -tags osusergo,netgo -o /application cmd/karate-server/main.go; 

FROM ubuntu:22.04

RUN apt-get update && apt-get install ca-certificates wget openjdk-17-jre git jq -y

RUN wget https://github.com/karatelabs/karate/releases/download/v1.2.0/karate-1.2.0.jar
RUN mv karate-1.2.0.jar karate.jar

COPY build/log-config.xml /

RUN apt-get install dnsutils -y
# DON'T CHANGE BELOW 
COPY --from=build /application /bin/application

EXPOSE 8080

CMD ["/bin/application", "--port=8080", "--host=0.0.0.0", "--write-timeout=0"]