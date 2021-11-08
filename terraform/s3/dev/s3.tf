module "s3" {
  source = "../_module/s3"

  service_name = "kakao-skillserver-s3"
  env_name = "dev"
}