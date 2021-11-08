module "logs" {
    source = "../_module/logs"

    env_name = "staging"
    crawling_lambda_role_name = data.terraform_remote_state.lambda.outputs.crawling_lambda_role_name
    menu_lambda_role_name = data.terraform_remote_state.lambda.outputs.menu_lambda_role_name
    skill_server_lambda_role_name = data.terraform_remote_state.lambda.outputs.skill_server_lambda_role_name
    skill_server_lambda_function_name = data.terraform_remote_state.lambda.outputs.skill_server_lambda_function_name
    menu_lambda_function_name = data.terraform_remote_state.lambda.outputs.menu_lambda_function_name
    crawling_lambda_function_name = data.terraform_remote_state.lambda.outputs.crawling_lambda_function_name
}