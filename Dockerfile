# Starting from the latest Golang container
FROM golang:1.5.1

# INSTALL any further tools you need here so they are cached in the docker build

# Set the WORKDIR to the project path in your GOPATH, e.g. /go/src/github.com/go-martini/martini/
WORKDIR /go/src/github.com/larskluge/babl/

# Copy the content of your repository into the container
COPY . ./

# Install dependencies through go get, unless you vendored them in your repository before
# Vendoring can be done through the godeps tool or Go vendoring available with
# Go versions after 1.5.1

RUN go get -t -v ./...

RUN go get github.com/mitchellh/gox
