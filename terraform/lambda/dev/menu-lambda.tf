module "menu-lambda" {
  source = "../_module/menu-lambda"

  s3_bucket_id = data.terraform_remote_state.s3.outputs.dh_bucket_id
  s3_bucket_arn = data.terraform_remote_state.s3.outputs.dh_bucket_arn
  env_name = "dev"
}