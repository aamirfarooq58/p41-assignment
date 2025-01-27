aws_region = "us-east-1"

lambda_config = {
  function_name = "app-function"
  runtime       = "provided.al2"
  handler       = "bootstrap"
  timeout       = 60
  memory_size   = 256
  zip_path      = "terraform/function/function.zip"
}

log_retention_days = 5
