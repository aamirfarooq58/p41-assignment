variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "vpc_config" {
  description = "VPC configuration"
  type = object({
    name               = string
    cidr               = string
    azs                = list(string)
    public_subnets     = list(string)
    private_subnets    = list(string)
    enable_nat_gateway = bool
    enable_vpn_gateway = bool
  })
  default = {
    name               = "app-vpc"
    cidr               = "10.0.0.0/16"
    azs                = ["us-east-1a", "us-east-1b"]
    public_subnets     = ["10.0.1.0/24", "10.0.2.0/24"]
    private_subnets    = ["10.0.3.0/24", "10.0.4.0/24"]
    enable_nat_gateway = true
    enable_vpn_gateway = true
  }
}

variable "lambda_config" {
  description = "Lambda function configuration"
  type = object({
    function_name = string
    runtime       = string
    handler       = string
    timeout       = number
    memory_size   = number
    zip_path      = string
  })
  default = {
    function_name = "app-function"
    runtime       = "provided.al2"
    handler       = "bootstrap"
    timeout       = 30
    memory_size   = 128
    zip_path      = "terraform/function/function.zip"
  }
}

variable "api_gateway_config" {
  description = "API Gateway configuration"
  type = object({
    name          = string
    protocol_type = string
  })
  default = {
    name          = "app-api"
    protocol_type = "HTTP"
  }
}

variable "log_retention_days" {
  description = "CloudWatch log retention in days"
  type        = number
  default     = 3
}
