# base image golang alpine
FROM golang:1.22-alpine AS builder

# set working directory
WORKDIR /app

# copy and build
COPY . .

RUN go build -o cutest .

# alpine image
FROM alpine:latest

COPY --from=builder /app/cutest /cutest
COPY --from=builder /app/index.html /index.html
# expose port
EXPOSE 8000 

# Define the command to run your application when the container starts
CMD ["./cutest"]
