output "skill_server_lambda_arn" {
  value = module.skill-server-lambda.skill_server_lambda_arn
}

output "skill_server_lambda_function_name" {
  value = module.skill-server-lambda.skill_server_lambda_function_name
}

output "skill_server_lambda_role_name" {
  value = module.skill-server-lambda.skill_server_lambda_role_name
}

output "menu_lambda_function_name" {
  value = module.menu-lambda.menu_lambda_function_name
}

output "menu_lambda_role_name" {
  value = module.menu-lambda.menu_lambda_role_name
}

output "crawling_lambda_function_name" {
  value = module.crawling-lambda.crawling_lambda_function_name
}

output "crawling_lambda_role_name" {
  value = module.crawling-lambda.crawling_lambda_role_name
}