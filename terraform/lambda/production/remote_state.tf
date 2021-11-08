data "terraform_remote_state" "s3" {
    backend = "s3"

    config = {
      bucket = "dh-prod-dh-apnortheast2-tfstate"
      region = "ap-northeast-2"
      key = "dh/terraform/s3/production/terraform.tfstate"
     }
}