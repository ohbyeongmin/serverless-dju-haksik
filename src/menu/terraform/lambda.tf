data "archive_file" "menu_function" {
    type = "zip"

    source_dir = "${path.module}/../bin"
    output_path = "${path.module}/menu.zip"
}

data "terraform_remote_state" "dh_s3" {
    backend = "s3"

    config = {
      bucket = "dh-tfstate-backend-extremely-able-buffalo"
      region = "ap-northeast-2"
      key = "crawling/terraform.tfstate"
     }
}

resource "aws_s3_bucket_object" "menu_function" {
  bucket = data.terraform_remote_state.dh_s3.outputs.dh_bucket_name

  key = "lambda_function/menu.zip"
  source = data.archive_file.menu_function.output_path

  etag = filemd5(data.archive_file.menu_function.output_path)
}

resource "aws_lambda_function" "menu_lambda" {
  function_name = "menu"

  s3_bucket = data.terraform_remote_state.dh_s3.outputs.dh_bucket_name
  s3_key = aws_s3_bucket_object.menu_function.key

  runtime = "go1.x"
  handler = "main"

  source_code_hash = data.archive_file.menu_function.output_base64sha256

  role = aws_iam_role.menu_lambda_role.arn

  environment {
    variables = {
      bucket = data.terraform_remote_state.dh_s3.outputs.dh_bucket_name
    }
  }
}

resource "aws_iam_role" "menu_lambda_role" {
  name = "menu_lambda"

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
  name        = "menu-policy"

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
  bucket = data.terraform_remote_state.dh_s3.outputs.dh_bucket_name
  

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
  source_arn    = data.terraform_remote_state.dh_s3.outputs.dh_bucket_arn
}