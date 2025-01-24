# Go Chat App

A simple chat application built with Go and WebSockets. It allows users to join chat rooms, send messages, and see active rooms and users.

This `README.md` now includes:

- **Key Files and Roles** for understanding the project structure.
- **Installation** and **Usage** instructions.
- **Endpoints** and **WebSocket message formats** for clarity.
- A **Deployment Guide** for optional deployment to AWS (including AWS App Runner, DynamoDB, S3, and Terraform for automation).


## Key Files and Their Roles

- **broadcasting.go**: Functions for broadcasting messages and user updates to all connected clients.
- **client.go**: Manages client connections and handles incoming messages from clients.
- **index.html**: Frontend of the chat application for interacting with users.
- **main.go**: Entry point of the application. It sets up HTTP routes and starts the server.
- **server.go**: Manages WebSocket connections, rooms, and message broadcasting.

## Installation

### Clone the repository:

```bash
git clone https://github.com/yourusername/go-chat-app.git
cd go-chat-app
```

### Install dependencies:

```bash
go mod tidy
```

## Usage

### Run the server:

```bash
go run main.go
```

### Open your browser and navigate to `http://localhost:8080` to access the chat application.

## Endpoints

- **/**: Serves the `index.html` file, the front end of the application.
- **/ws**: Handles WebSocket connections for real-time communication.
- **/active-rooms**: Returns a list of active chat rooms.
- **/room-users**: Returns a list of users in a specified room.

## WebSocket Messages

### Message Structure

When sending messages via WebSocket, the structure of each message is as follows:

```json
{
    "type": "message",
    "content": "Hello, World!",
    "sender": "username",
    "room": "roomName"
}
```

### Message Types

- **join**: A user joins a room.
- **leave**: A user leaves a room.
- **message**: A user sends a message to the chat room.
- **switch**: A user switches from one room to another.

## Deployment Guide (Optional for Showcase)

This section outlines how to deploy the application using AWS services like AWS App Runner and DynamoDB for storing messages.

### Step 1: Set Up AWS App Runner for Deployment

AWS App Runner is a serverless container service that makes it easy to deploy Go applications.

- Build and push your Docker image to Amazon Elastic Container Registry (ECR).
  
```bash
aws ecr create-repository --repository-name chat-app
# Authenticate Docker to the ECR registry
aws ecr get-login-password --region <region> | docker login --username AWS --password-stdin <aws_account_id>.dkr.ecr.<region>.amazonaws.com
docker tag <image_name> <aws_account_id>.dkr.ecr.<region>.amazonaws.com/chat-app:latest
docker push <aws_account_id>.dkr.ecr.<region>.amazonaws.com/chat-app:latest
```

- Create an App Runner service that points to the image in ECR.
- Configure auto-scaling, networking, and environment variables as required.

### Step 2: Use DynamoDB for Storing Messages

To store messages in a serverless, scalable database, use AWS DynamoDB.

- Create a table in DynamoDB to store messages. Example schema:
    - Table name: `ChatMessages`
    - Partition key: `roomId`
    - Sort key: `timestamp`

- In the Go app, update the WebSocket handler to store and retrieve messages from DynamoDB.

### Step 3: Optional S3 Integration for File Uploads

If you want to add file upload capabilities (e.g., for sharing images or documents in the chat), integrate AWS S3 for file storage.

- Set up an S3 bucket in AWS.
- Modify the frontend to allow users to upload files.
- Store the uploaded files in the S3 bucket and share the URLs in the chat.

### Step 4: Automate Deployment with Terraform

If you'd like to showcase your skills in infrastructure automation, use **Terraform** to provision the necessary AWS resources such as:

- AWS App Runner for container deployment
- DynamoDB for message storage
- S3 for file uploads

Here’s an example of how you might set up Terraform for AWS services:

```hcl
# Terraform AWS provider configuration
provider "aws" {
  region = "us-east-1"
}

# AWS App Runner service
resource "aws_apprunner_service" "chat_app" {
  service_name = "chat-app-service"
  source_configuration {
    image_repository {
      image_identifier = "<aws_account_id>.dkr.ecr.<region>.amazonaws.com/chat-app:latest"
      image_configuration {
        port = "8080"
      }
    }
  }
}

# DynamoDB Table for messages
resource "aws_dynamodb_table" "chat_messages" {
  name         = "ChatMessages"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "roomId"
  range_key    = "timestamp"
  attribute {
    name = "roomId"
    type = "S"
  }
  attribute {
    name = "timestamp"
    type = "S"
  }
}

# S3 bucket for file uploads (optional)
resource "aws_s3_bucket" "chat_files" {
  bucket = "chat-app-files"
}
```

### Step 5: Clean Up After Deployment

Remember to clean up resources after you’re done with the deployment to avoid incurring unnecessary charges:

```bash
terraform destroy  # Deletes all resources created by Terraform
```

## Note

This project is developed primarily for the purpose of learning, practicing, and showcasing skills in cloud computing, Go development, and Terraform. It is not intended for production deployment or commercial use. The infrastructure and configuration presented here are meant to demonstrate the developer's understanding and proficiency in these technologies.
```
