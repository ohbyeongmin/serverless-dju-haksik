data "terraform_remote_state" "lambda" {
    backend = "s3"

    config = {
      bucket = "dh-prod-dh-apnortheast2-tfstate"
      region = "ap-northeast-2"
      key = "dh/terraform/lambda/production/terraform.tfstate"
     }
}