output "dh_bucket_name" {
    description = "Name of the S3 bucket used to dh kakao bot"

    value = aws_s3_bucket.dh_bucket.id
}

output "dh_crawling_lambda" {
    description = "Name of the lambda function used to crawling"

    value = aws_lambda_function.crawling_lambda.function_name
}

output "dh_bucket_arn" {
    value = aws_s3_bucket.dh_bucket.arn
}