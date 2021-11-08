data "archive_file" "menu_function_zip" {
    type = "zip"

    source_dir = "${path.module}/../../${var.env_name}/function/menu/bin"
    output_path = "${path.module}/../../${var.env_name}/function/menu/zip/menu.zip"
}

resource "aws_s3_bucket_object" "menu_function" {
  bucket = "${var.s3_bucket_id}"

  key = "lambda_function/menu.zip"
  source = data.archive_file.menu_function_zip.output_path

  etag = filemd5(data.archive_file.menu_function_zip.output_path)
}

resource "aws_lambda_function" "menu_lambda" {
  function_name = "menu-${var.env_name}"

  s3_bucket = "${var.s3_bucket_id}"
  s3_key = aws_s3_bucket_object.menu_function.key

  runtime = "go1.x"
  handler = "main"

  source_code_hash = data.archive_file.menu_function_zip.output_base64sha256

  role = aws_iam_role.menu_lambda_role.arn

  environment {
    variables = {
      bucket = "${var.s3_bucket_id}"
    }
  }
}

resource "aws_iam_role" "menu_lambda_role" {
  name = "menu-lambda-${var.env_name}"

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
  name        = "menu-policy-${var.env_name}"

  policy = jsonencode({
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "lambda:InvokeFunction",
        "s3:PutObject",
        "s3:GetObject"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
  })
}

resource "aws_iam_role_policy_attachment" "menu_lambda_policy" {
  role       = aws_iam_role.menu_lambda_role.name
  policy_arn = aws_iam_policy.policy.arn
}


resource "aws_s3_bucket_notification" "bucket_notification" {
  bucket = "${var.s3_bucket_id}"
  

  lambda_function {
    lambda_function_arn = aws_lambda_function.menu_lambda.arn 
    events = ["s3:ObjectCreated:Put"]
    filter_prefix = "data"
    filter_suffix = ".xlsx"
  }

  depends_on = [
     aws_lambda_permission.allow_bucket
  ]
}

resource "aws_lambda_permission" "allow_bucket" {
  statement_id = "AllowExecutionFromS3Bucket"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.menu_lambda.arn
  principal     = "s3.amazonaws.com"
  source_arn    = "${var.s3_bucket_arn}"
}