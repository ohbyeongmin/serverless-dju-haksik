output "dh_bucket_name" {
    description = "Name of the S3 bucket used to dh kakao bot"

    value = aws_s3_bucket.dh_bucket.id
}