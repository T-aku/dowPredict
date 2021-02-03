FROM golang:latest
RUN mkdir /golang && go get -u github.com/gorilla/websocket && go get -u github.com/lib/pq