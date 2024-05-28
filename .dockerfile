# base image golang alpine
FROM golang:1.22-alpine AS builder

# set working directory
WORKDIR /app

# copy and build
COPY . .
RUN go build -o main .

# expose port
EXPOSE 8000 

# Define the command to run your application when the container starts
CMD ["./cutest"]
