 Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.11 as builder
WORKDIR /go/src/github.com/space/project/

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/space/project/

# Build the service inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM centurylink/ca-certs
EXPOSE 8080

# Copy app
COPY --from=builder /go/src/github.com/space/project/app   /
ENTRYPOINT ["/app"]