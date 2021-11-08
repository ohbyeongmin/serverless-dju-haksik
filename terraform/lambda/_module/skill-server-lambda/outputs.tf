output "skill_server_lambda_arn" {
  value = aws_lambda_function.skill_server_lambda.arn
}

output "skill_server_lambda_function_name" {
  value = aws_lambda_function.skill_server_lambda.function_name
}

output "skill_server_lambda_role_name" {
  value = aws_iam_role.skill_server_lambda_role.name
}