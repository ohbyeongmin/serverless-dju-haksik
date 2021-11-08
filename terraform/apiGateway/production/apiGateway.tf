module "apiGateway" {
    source = "../_module/apiGateway"

    env_name = "production"
    skill_server_lambda_arn = data.terraform_remote_state.lambda.outputs.skill_server_lambda_arn
    skill_server_lambda_function_name = data.terraform_remote_state.lambda.outputs.skill_server_lambda_function_name
}