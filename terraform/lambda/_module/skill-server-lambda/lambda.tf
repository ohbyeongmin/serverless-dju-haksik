data "archive_file" "skill_server_function_zip" {
    type = "zip"


    source_dir = "${path.module}/../../${var.env_name}/function/skill-server/bin"
    output_path = "${path.module}/../../${var.env_name}/function/skill-server/zip/skillserver.zip"
}


resource "aws_s3_bucket_object" "skill_server_function" {
  bucket = "${var.s3_bucket_id}"

  key = "lambda_function/skillserver.zip"
  source = data.archive_file.skill_server_function_zip.output_path

  etag = filemd5(data.archive_file.skill_server_function_zip.output_path)
}



resource "aws_lambda_function" "skill_server_lambda" {
  function_name = "skillServer-${var.env_name}"

  s3_bucket = "${var.s3_bucket_id}"
  s3_key = aws_s3_bucket_object.skill_server_function.key

  runtime = "go1.x"
  handler = "main"

  source_code_hash = data.archive_file.skill_server_function_zip.output_base64sha256

  role = aws_iam_role.skill_server_lambda_role.arn

  environment {
    variables = {
      bucket = "${var.s3_bucket_id}"
      objectFile = "${var.object_file_name}"
    }
  }
}


resource "aws_iam_role" "skill_server_lambda_role" {
  name = "skill-server-lambda-${var.env_name}"

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
  name        = "skill-server-policy-${var.env_name}"

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

