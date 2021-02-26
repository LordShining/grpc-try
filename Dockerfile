FROM golang:1.15

RUN cd /go/src \
&& git clone https://github.com/LordShining/MyFirstGolangWebserver.git \
&& cd MyFirstGolangWebserver \
&& go get -d -v \
&& go install -v
WORKDIR /go/src/MyFirstGolangWebserver

ENTRYPOINT [ "go", "run", "/go/src/MyFirstGolangWebserver/main.go" ]