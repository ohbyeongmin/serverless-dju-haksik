terraform {
  required_version = "= 1.0.9" # Terraform Version 

  backend "s3" {
    bucket         = "dh-prod-dh-apnortheast2-tfstate" # Set bucket name 
    key            = "dh/terraform/ec2/dh-prod/terraform.tfstate"
    region         = "ap-northeast-2"
    encrypt        = true
    dynamodb_table = "dh-prod-terraform-lock" # Set DynamoDB Table
  }
}