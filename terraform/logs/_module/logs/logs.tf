resource "aws_iam_policy" "logs_policy" {
  name ="logs-policy-${var.env_name}"

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

resource "aws_iam_role_policy_attachment" "logs_crawling_lambda_policy" {
  role       = "${var.crawling_lambda_role_name}"
  policy_arn = aws_iam_policy.logs_policy.arn
}

resource "aws_iam_role_policy_attachment" "logs_menu_lambda_policy" {
  role       = "${var.menu_lambda_role_name}"
  policy_arn = aws_iam_policy.logs_policy.arn
}

resource "aws_iam_role_policy_attachment" "logs_skill_server_lambda_policy" {
  role       = "${var.skill_server_lambda_role_name}"
  policy_arn = aws_iam_policy.logs_policy.arn
}

resource "aws_cloudwatch_log_group" "skill_server_function" {
  name = "/aws/lambda/${var.skill_server_lambda_function_name}"

  retention_in_days = 180
}

resource "aws_cloudwatch_log_group" "menu_function" {
  name = "/aws/lambda/${var.menu_lambda_function_name}"

  retention_in_days = 180
}

resource "aws_cloudwatch_log_group" "crawling_function" {
  name = "/aws/lambda/${var.crawling_lambda_function_name}"

  retention_in_days = 180
}