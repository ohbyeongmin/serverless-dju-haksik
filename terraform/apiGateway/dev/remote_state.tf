data "terraform_remote_state" "lambda" {
    backend = "s3"

    config = {
      bucket = "byeongmin-dh-apnortheast2-tfstate"
      region = "ap-northeast-2"
      key = "dh/terraform/lambda/dev/terraform.tfstate"
     }
}