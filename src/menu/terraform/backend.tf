terraform {
    backend "s3" {  
      bucket         = "dh-tfstate-backend-extremely-able-buffalo"
      key            = "menu/terraform.tfstate"
      region         = "ap-northeast-2"
      encrypt        = true
      dynamodb_table = "dh-terraform-lock"
    }
}