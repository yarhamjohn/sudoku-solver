# First stage: start with a Golang base image
FROM golang:1.12-alpine3.10

# Move to the directory where the source code will live
WORKDIR /go/src/sudoku-solver

# Copy the source code into the current directory
COPY ./Golang .

# Get any dependencies, and compile the code
RUN CGO_ENABLED=0 go get -v ./...

# Second stage: start from an empty base image
FROM scratch

# Copy the binary from the first stage
COPY --from=0 /go/bin/sudoku-solver /

# Tell Docker what executable to run by default when starting this container
ENTRYPOINT ["/sudoku-solver"]