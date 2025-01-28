# Serverless Application on AWS

# Prerequisites
1. AWS CLI installed and configured

2. Go 1.x or later

3. Terraform 1.x or later

4. Valid AWS credentials set up

# Deployment
```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap function/main.go
zip function.zip function/bootstrap

# Deploy infrastructure
cd serverless
export AWS_PROFILE = <your-profile>
terraform init
terraform plan
terraform apply

## Cleanup
terraform destroy

## Set Concourse Pipeline
Update the vars.yaml file.
cd terraform/Pipeline
fly -t your-target set-pipeline -p lambda-deploy -c pipeline.yml -l vars.yml
