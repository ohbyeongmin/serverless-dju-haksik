output "menu_lambda_function_name" {
  value = aws_lambda_function.menu_lambda.function_name
}

output "menu_lambda_role_name" {
  value = aws_iam_role.menu_lambda_role.name
}