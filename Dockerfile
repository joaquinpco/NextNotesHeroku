FROM golang
WORKDIR /go/src/github.com/nextnotes
COPY . ./
RUN go install .
CMD ["/go/bin/nextnotes"]