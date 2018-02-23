FROM golang

ADD . /go/src/github.com/lordalek/milage-api

WORKDIR /go/src/github.com/lordalek/milage-api
RUN go get github.com/gorilla/mux && go get github.com/nu7hatch/gouuid

RUN go install github.com/lordalek/milage-api


# Expose port 8080 to the host so that outer-world can access your application
EXPOSE 8080

ENTRYPOINT /go/bin/milage-api
