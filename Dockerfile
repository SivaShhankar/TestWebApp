# golang image where workspace (GOPATH) configured at /go.
FROM golang:latest

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/SivaShhankar/TestWebApp

# Build the golang-docker command inside the container.
RUN go install github.com/SivaShhankar/TestWebApp

# Run the golang-docker command by default when the container starts.
ENTRYPOINT /go/bin/TestWebApp

# http server listens on port 8080.
EXPOSE 8080