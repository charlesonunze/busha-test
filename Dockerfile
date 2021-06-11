FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/charlesonunze/busha-test

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 3000 6379

# RUN ["chmod", "+x", "/usr/bin/busha-test"]

# Run the executable
CMD ["busha-test"]