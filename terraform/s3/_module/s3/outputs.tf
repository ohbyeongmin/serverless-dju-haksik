output "dh_bucket_id" {
    description = "id of the S3 bucket used to dh kakao bot"

    value = aws_s3_bucket.s3.id
}