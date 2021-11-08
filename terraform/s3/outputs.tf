output "dh_bucket_id" {
    description = "id of the S3 bucket used to dh kakao bot"

    value = module.s3.dh_bucket_id
}

output "dh_bucket_arn" {
    description = "arn of the S3 bucket used to dh kakao bot"

    value = module.s3.dh_bucket_arn
}