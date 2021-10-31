data "archive_file" "skill_server_function" {
    type = "zip"

    source_dir = "${path.module}/../bin"
    output_path = "${path.module}/skillserver.zip"
}

data "terraform_remote_state" "dh_s3" {
    backend = "s3"

    config = {
      bucket = "dh-tfstate-backend-extremely-able-buffalo"
      region = "ap-northeast-2"
      key = "crawling/terraform.tfstate"
     }
}

data "terraform_remote_state" "object_file" {
  backend = "s3"

  config = {
    bucket = "dh-tfstate-backend-extremely-able-buffalo"
      region = "ap-northeast-2"
      key = "menu/terraform.tfstate"
  }
}

resource "aws_s3_bucket_object" "skill_server_function" {
  bucket = data.terraform_remote_state.dh_s3.outputs.dh_bucket_name

  key = "lambda_function/skillserver.zip"
  source = data.archive_file.skill_server_function.output_path

  etag = filemd5(data.archive_file.skill_server_function.output_path)
}



resource "aws_lambda_function" "skill_server_lambda" {
  function_name = "skillServer"

  s3_bucket = data.terraform_remote_state.dh_s3.outputs.dh_bucket_name
  s3_key = aws_s3_bucket_object.skill_server_function.key

  runtime = "go1.x"
  handler = "main"

  source_code_hash = data.archive_file.skill_server_function.output_base64sha256

  role = aws_iam_role.skill_server_lambda_role.arn

  environment {
    variables = {
      bucket = data.terraform_remote_state.dh_s3.outputs.dh_bucket_name
      objectFile = data.terraform_remote_state.object_file.outputs.object_file_name
    }
  }
}

resource "aws_cloudwatch_log_group" "skill_server_function" {
  name = "/aws/lambda/${aws_lambda_function.skill_server_lambda.function_name}"

  retention_in_days = 30
}

resource "aws_iam_role" "skill_server_lambda_role" {
  name = "skill_server_lambda"

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
  name        = "skill-server-policy"

  policy = jsonencode({
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "lambda:InvokeFunction",
        "s3:GetObject"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
  })
}

resource "aws_iam_role_policy_attachment" "skill_server_lambda_policy" {
  role       = aws_iam_role.skill_server_lambda_role.name
  policy_arn = aws_iam_policy.policy.arn
}

