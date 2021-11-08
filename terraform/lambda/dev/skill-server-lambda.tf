module "skill-server-lambda" {
    source = "../_module/skill-server-lambda"

  s3_bucket_id = data.terraform_remote_state.s3.outputs.dh_bucket_id
  env_name = "dev"
  object_file_name = "menuObject"
}