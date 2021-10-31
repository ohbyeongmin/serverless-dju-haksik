resource "aws_apigatewayv2_api" "skill_server_lambda" {
  name          = "skill_server_lambda_gw"
  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_stage" "skill_server_lambda" {
  api_id = aws_apigatewayv2_api.skill_server_lambda.id

  name        = "skill_server_lambda_stage"
  auto_deploy = true

  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.api_gw.arn

    format = jsonencode({
      requestId               = "$context.requestId"
      sourceIp                = "$context.identity.sourceIp"
      requestTime             = "$context.requestTime"
      protocol                = "$context.protocol"
      httpMethod              = "$context.httpMethod"
      resourcePath            = "$context.resourcePath"
      routeKey                = "$context.routeKey"
      status                  = "$context.status"
      responseLength          = "$context.responseLength"
      integrationErrorMessage = "$context.integrationErrorMessage"
      }
    )
  }
}


// API Gateway --> Lambda
resource "aws_apigatewayv2_integration" "menu" {
  api_id = aws_apigatewayv2_api.skill_server_lambda.id

  integration_uri    = aws_lambda_function.skill_server_lambda.arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}

// internet --> API Gateway 
resource "aws_apigatewayv2_route" "menu" {
  api_id = aws_apigatewayv2_api.skill_server_lambda.id

  route_key = "POST /menu"
  target    = "integrations/${aws_apigatewayv2_integration.menu.id}"
}

resource "aws_cloudwatch_log_group" "api_gw" {
  name = "/aws/api_gw/${aws_apigatewayv2_api.skill_server_lambda.name}"

  retention_in_days = 30
}

resource "aws_lambda_permission" "api_gw" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.skill_server_lambda.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.skill_server_lambda.execution_arn}/*/*"
}