# AWS, Go REST API with S3 File Storage

A REST API built with Go and Gin, deployed on AWS EC2 with S3 file storage. Built as an intro project to learn core AWS services.

## Stack
- **Go + Gin** — REST API
- **AWS EC2** (t3.small) — server, managed by systemd for 24/7 uptime
- **AWS S3** — file storage with presigned URLs (15 min expiry)
- **AWS IAM** — EC2 instance role for secure S3 access, no hardcoded credentials

## Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/test` | Health check |
| POST | `/api/v1/upload` | Upload a file to S3 |
| GET | `/api/v1/files` | List all uploaded files with presigned URLs |

## Project Structure

```
├── main.go
├── config/        # AWS S3 client + constants
├── handlers/      # Upload, list, and test handlers
├── middleware/    # CORS
└── routes/        # Route registration
```

## Deployment

Compiled locally for Linux and uploaded to EC2 as a binary — no compilation on the server.

```bash
# Build for Linux
GOOS=linux GOARCH=amd64 go build -o app

# Upload to EC2
scp -i your-key.pem ./app ubuntu@<ec2-ip>:~/AWS-learn/

# Restart the service
sudo systemctl restart goapi
```

## What I Learned
- Launching and configuring an EC2 instance
- Setting up security groups and inbound rules
- IAM roles for secure service-to-service auth
- S3 file uploads and presigned URL generation
- Keeping a Go server running 24/7 with systemd
