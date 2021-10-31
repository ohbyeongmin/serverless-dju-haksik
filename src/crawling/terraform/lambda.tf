data "archive_file" "crawling_function" {
    type = "zip"

    source_dir = "${path.module}/../bin"
    output_path = "${path.module}/crawling.zip"
}

resource "aws_s3_bucket_object" "crawling_function" {
  bucket = aws_s3_bucket.dh_bucket.id

  key = "lambda_function/crawling.zip"
  source = data.archive_file.crawling_function.output_path

  etag = filemd5(data.archive_file.crawling_function.output_path)
}

resource "aws_lambda_function" "crawling_lambda" {
  function_name = "crawling"

  s3_bucket = aws_s3_bucket.dh_bucket.id
  s3_key = aws_s3_bucket_object.crawling_function.key

  runtime = "go1.x"
  handler = "main"

  source_code_hash = data.archive_file.crawling_function.output_base64sha256

  role = aws_iam_role.crawling_lambda_role.arn

  environment {
    variables = {
      bucket = aws_s3_bucket.dh_bucket.id
    }
  }
}

resource "aws_iam_role" "crawling_lambda_role" {
  name = "crawling_lambda"

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
  name        = "crawling-policy"

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

resource "aws_iam_policy" "logs_policy" {
  name ="logs-policy"

  policy = jsonencode({
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:logs:*:*:*"
    }
  ]
  })
}


resource "aws_iam_role_policy_attachment" "crawling_lambda_policy" {
  role       = aws_iam_role.crawling_lambda_role.name
  policy_arn = aws_iam_policy.policy.arn
}

resource "aws_iam_role_policy_attachment" "logs_lambda_policy" {
  role       = aws_iam_role.crawling_lambda_role.name
  policy_arn = aws_iam_policy.logs_policy.arn
}
