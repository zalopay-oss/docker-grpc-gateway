############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add git
# Create appuser.
RUN adduser -D -g '' appuser
# WORKDIR $GOPATH/src/mypackage/myapp/
WORKDIR /myapp
ENV GO111MODULE=on
COPY go.mod go.sum ./

RUN go mod download

COPY . /myapp

RUN rm -rf ~/.netrc

# Build the binary.
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build -ldflags="-w -s" -o /go/bin/myapp cmd/server/main.go 
############################
# STEP 2 build a small image
############################
FROM scratch
# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd

WORKDIR /myapp
# Copy our static executable.
COPY --from=builder /go/bin/myapp ./


# Use an unprivileged user.
# USER appuser

# Run the hello binary. 
EXPOSE 8000
ENTRYPOINT ["./myapp"]
