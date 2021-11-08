module "crawling-lambda" {
  source = "../_module/crawling-lambda"

  s3_bucket_id = data.terraform_remote_state.s3.outputs.dh_bucket_id
  env_name = "staging"
  schedule_expression_cron = "cron(0 2 ? * MON *)"
}