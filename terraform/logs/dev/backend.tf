terraform {
  required_version = "= 1.0.9" # Terraform Version 

  backend "s3" {
    bucket         = "byeongmin-dh-apnortheast2-tfstate" # Set bucket name 
    key            = "dh/terraform/logs/dev/terraform.tfstate"
    region         = "ap-northeast-2"
    encrypt        = true
    dynamodb_table = "bm-dh-terraform-lock" # Set DynamoDB Table
  }
}