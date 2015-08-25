#MAINTAINER Mikhail Yarotski playaer80@gmail.com

# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM playaer/my-test

ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin


ENV APP_DOMAIN http://localhost:3000

ENV APP_EMAIL_USER playaer.my.test@gmail.com
ENV APP_EMAIL_PASS mkmk1980
ENV APP_EMAIL_HOST smtp.gmail.com
ENV APP_EMAIL_PORT 587

ENV APP_DB_NAME first_go
ENV APP_DB_USER root
ENV APP_DB_PASS 111


ADD . /go/src/github.com/playaer/myFirstGoProject

RUN go get github.com/go-martini/martini && go get github.com/go-sql-driver/mysql && go get github.com/martini-contrib/render
RUN go install github.com/playaer/myFirstGoProject

# Run the outyet command by default when the container starts.
ENTRYPOINT cd /go/src/github.com/playaer/myFirstGoProject && service mysql start && /go/bin/myFirstGoProject && bash

# Document that the service listens on port 8080.
EXPOSE 3000




