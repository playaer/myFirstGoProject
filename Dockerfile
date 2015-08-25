# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/playaer/myFirstGoProject

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get github.com/go-martini/martini && go get github.com/go-sql-driver/mysql && go get github.com/martini-contrib/render && go install github.com/playaer/myFirstGoProject

#RUN DEBIAN_FRONTEND=noninteractive apt-get install -y mysql-server

# Run the outyet command by default when the container starts.
ENTRYPOINT /bin/bash #/go/bin/myFirstGoProject

# Document that the service listens on port 8080.
EXPOSE 3000
