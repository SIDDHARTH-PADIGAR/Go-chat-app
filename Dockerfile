FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy all files from the current directory to the working directory
COPY . .

# Install Go modules and build the app
RUN go mod tidy && go build -o chat-app

#Expose the port the app will run on
EXPOSE 8080

# Run the app
CMD ["./chat-app"]
