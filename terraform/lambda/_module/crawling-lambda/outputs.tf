output "crawling_lambda_function_name" {
  value = aws_lambda_function.crawling_lambda.function_name 
}

output "crawling_lambda_role_name" {
  value = aws_iam_role.crawling_lambda_role.name 
}