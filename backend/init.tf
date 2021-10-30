terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.48.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.1.0"
    }
  }
}

provider "aws" {
  region = var.aws-region
}

resource "random_pet" "tfstate_bucket_name" {
  prefix = "dh-tfstate-backend"
  length = 3
}

resource "aws_s3_bucket" "tfstate" {
  bucket = random_pet.tfstate_bucket_name.id

  versioning {
    enabled = true # Prevent from deleting tfstate file
  }
}

#DynamoDB for terraform state lock
resource "aws_dynamodb_table" "terraform_state_lock" {
  name = "dh-terraform-lock"
  hash_key = "LockID"
  billing_mode   = "PAY_PER_REQUEST"

  attribute {
    name = "LockID"
    type = "S"
  }
}