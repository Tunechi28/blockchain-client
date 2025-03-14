variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-west-2"
}

variable "vpc_cidr" {
  description = "CIDR block for the VPC"
  type        = string
  default     = "10.0.0.0/16"
}

variable "public_subnet_cidrs" {
  description = "CIDR blocks for public subnets"
  type        = list(string)
  default     = ["10.0.1.0/24", "10.0.2.0/24"]
}

variable "ecs_cluster_name" {
  description = "ECS cluster name"
  type        = string
  default     = "blockchain-client-cluster"
}

variable "services" {
  description = "Map of services to deploy"
  type = map(object({
    container_name = string
    image          = string
    container_port = number
    cpu            = string
    memory         = string
    desired_count  = number
  }))
  default = {
    "blockchain-client-service-1" = {
      container_name = "blockchain-client"
      image          = "blockchain-client:latest"
      container_port = 8080
      cpu            = "256"
      memory         = "512"
      desired_count  = 1
    }
  }
}
