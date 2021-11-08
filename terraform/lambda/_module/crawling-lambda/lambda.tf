data "archive_file" "crawling_function_zip" {
    type = "zip"

    source_dir = "${path.module}/../../${var.env_name}/function/crawling/bin"
    output_path = "${path.module}/../../${var.env_name}/function/crawling/zip/crawling.zip"
}

resource "aws_s3_bucket_object" "crawling_function" {
  bucket = "${var.s3_bucket_id}"

  key = "lambda_function/crawling.zip"
  source = data.archive_file.crawling_function_zip.output_path

  etag = filemd5(data.archive_file.crawling_function_zip.output_path)
}

resource "aws_lambda_function" "crawling_lambda" {
  function_name = "crawling-${var.env_name}"

  s3_bucket = "${var.s3_bucket_id}"
  s3_key = aws_s3_bucket_object.crawling_function.key

  runtime = "go1.x"
  handler = "main"

  source_code_hash = data.archive_file.crawling_function_zip.output_base64sha256

  role = aws_iam_role.crawling_lambda_role.arn

  environment {
    variables = {
      bucket = "${var.s3_bucket_id}"
    }
  }
}

resource "aws_iam_role" "crawling_lambda_role" {
  name = "crawling-lambda-${var.env_name}"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Sid    = ""
      Principal = {
        Service = "lambda.amazonaws.com"
      }
      }
    ]
  })
}

resource "aws_iam_policy" "policy" {
  name        = "crawling-policy-${var.env_name}"

  policy = jsonencode({
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "lambda:InvokeFunction",
        "s3:PutObject"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
  })
}

resource "aws_iam_role_policy_attachment" "crawling_lambda_policy" {
  role       = aws_iam_role.crawling_lambda_role.name
  policy_arn = aws_iam_policy.policy.arn
}