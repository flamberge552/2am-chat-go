FROM golang:1.12

# enable go mods
ENV GO111MODULE=on

# set workdir
WORKDIR $GOPATH/src/github.com/gabrielciordas/2am-chat-go

# copy my stuff into container instance
COPY . .

# install deps
RUN go get -d -v ./...
RUN go install -v ./...

# the port we want our app to run on, at some point this will be passed in as an env var
EXPOSE 8080

# this is the docker equivalent of a double click on an executable (for this specific scenario of course)
CMD ["2am-chat-go"]