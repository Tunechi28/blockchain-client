# Blockchain Client

## Overview
This project is a simple blockchain client implemented in Go. It interacts with the Polygon network by fetching the latest block number and retrieving block details by number using JSON-RPC. The API is exposed via REST endpoints using the Gin framework.

## 🚀 Features

- 🌐 **REST API** with `Gin` framework
- 🔗 **Blockchain RPC Client** (Fetches latest block number and block details)
- ✅ **Unit Tests** to ensure reliability
- 📦 **Dockerized** for seamless deployment
- 🌍 **Terraform Scripts** for AWS ECS Fargate deployment
- 📜 **Swagger Documentation** for API endpoints

## 🏗 Architecture

### 🖥 **1. API Server (`main.go`)**
- Serves **HTTP endpoints** using **Gin**

### 🔗 **2. Blockchain RPC Client (`client/`)**
- Handles API requests to **Polygon RPC**
- Implements:
  - 🏗 **Get latest block number** (`eth_blockNumber`)
  - 🔍 **Get block by number** (`eth_getBlockByNumber`)

### 🏗 **3. Handlers (`handlers/`)**
- Defines **HTTP request handlers** for API endpoints
- Uses **Gin Context** for request handling

### 📜 **4. Swagger API Documentation**
- **Auto-generates OpenAPI docs**
- Available at **`/swagger/index.html`**

### 🐳 **5. Docker & Terraform**
- **Dockerized** for portability
- **Terraform HCL** for AWS ECS deployment

---

## 🛠 Tech Stack

- **Programming Language:** Go
- **Web Framework:** [Gin](https://gin-gonic.com/)
- **Blockchain Network:** Polygon
- **Infrastructure:** Docker, Terraform, AWS ECS Fargate
- **Documentation:** Swagger

---

## 🔗 API Endpoints

| Endpoint                     | Method | Description               |
|------------------------------|--------|---------------------------|
| `/api/block/latest`          | GET    | Fetch latest block number |
| `/api/block/:blockNumber`    | GET    | Fetch block details       |


## 🚀 Getting Started
### Prerequisites
- Go 1.22.1 or later
- Docker (for containerization)
- Terraform (for infrastructure as code)

### Running Locally

1. **Clone the repository:**
   ```bash
   git clone https://github.com/tunechi28/blockchain-client.git
   cd blockchain-client
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the application:**
   ```bash
   go run main.go
   ```

   The application will start on `http://localhost:8080`.

---

### Testing Application
1. **Run unit tests:**
   ```bash
   go test ./... -v
   ```

2. **View test coverage:**
   ```bash
   go test ./... -cover
   ```

### Building and Running with Docker

1. **Build the Docker image:**
   ```bash
   docker build -t blockchain-client .
   ```

2. **Run the Docker container:**
   ```bash
   docker run -p 8080:8080 blockchain-client
   ```

   The application will be accessible at `http://localhost:8080`.

---

### Accessing Swagger Documentation

Once the application is running, access the Swagger UI at:
```
http://localhost:8080/swagger/index.html
```

---

### Terraform Deployment (Local State)

The Terraform configuration in the `/terraform` folder can be used to deploy the application to AWS ECS Fargate. Note that this setup uses **local state** and does not require actual deployment.

1. **Navigate to the Terraform folder:**
   ```bash
   cd terraform
   ```

2. **Initialize Terraform:**
   ```bash
   terraform init
   ```

3. **Review the Terraform plan:**
   ```bash
   terraform plan
   ```
   This will show the resources that Terraform will create. No actual deployment is required for this exercise.

---

## 🔥 Production Ready Requirement

### 🔐 Security Enhancements
- Implement **API authentication** using JWT or API keys to restrict access.
- Enforce **rate limiting** to prevent abuse and mitigate DDoS attacks.
- Require **HTTPS** for secure communication.

### 📊 Observability
- Use **structured logging** with `logrus` or `zap` to improve log readability and debugging.
- Integrate **Prometheus and Grafana** for real-time monitoring and performance insights.
- Enable **distributed tracing** using OpenTelemetry to track request flows and diagnose latency issues.

### ⚡ Resilience & Reliability
- Implement **automatic retries and exponential backoff** for failed blockchain RPC calls.
- Use **circuit breakers** to gracefully handle provider failures.
- Enable **error tracking** with tools like Sentry or Rollbar.

### 🏗 Scalability & Performance
- Deploy **auto-scaling ECS tasks** to handle variable loads efficiently.
- Optimize API performance by **caching blockchain responses** in Redis.
- Reduce latency by **optimizing API calls** and implementing efficient request batching.

### 🛠 CI/CD & Deployment
- Automate **deployment** with GitHub Actions, it currently only supports running unit tests.
- Perform **vulnerability scanning** on Docker images using Trivy.
- Implement **integration and load testing** to ensure stability under real-world conditions.

### 📁 Configuration Management
- Use **environment variables** to manage sensitive configurations securely.
- Implement **secret management** using AWS Secrets Manager or HashiCorp Vault.

### 🌍 Remote Terraform State
- Store Terraform state in **S3 with DynamoDB for state locking** to support multi-developer collaboration.
- Implement **advanced networking and security configurations** in Terraform for production-ready infrastructure.

### 🔄 Functionality Enhancements
- **Multi-Provider Support**: Integrate multiple blockchain RPC providers (e.g., **Alchemy, Infura, QuickNode, Chainstack**) for higher reliability.
- **Automatic Provider Switching**: Implement a failover mechanism to detect provider failures and switch automatically.
- **Smart Proxy**: Route API requests through a reverse proxy with load balancing, caching, and rate limiting.
- **Latency-Based Routing**: Dynamically choose the fastest available provider for optimal performance.
- **Batch Requests & Historical Data Storage**: Support bulk blockchain queries and store relevant data in **PostgreSQL** or **ClickHouse** for improved querying.
