### BASE GO IMAGE
FROM golang:1.22-alpine as builder

# creates app folder
RUN mkdir /app

# copies everything from the root folder into app folder
COPY ./ /app

# sets app folder as working directory for next instructions
WORKDIR /app

# builds the application executable into the previously set working directory (app folder)
RUN CGO_ENABLED=0 go build -o application .

# adds the executable rights to application executable
RUN chmod +x /app/application



### BUILD A TINY DOCKER IMAGE
FROM alpine:latest

# creates app folder
RUN mkdir /app

# copies the previously created builder image from /app/application of the previous image into /app of the current image 
COPY --from=builder /app/application /app

# runs the executable
CMD ["/app/application"]