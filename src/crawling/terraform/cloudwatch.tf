resource "aws_cloudwatch_event_rule" "every_monday_11AM" {
  name = "every_monday_11AM"
  description = "Fires every monday 11 AM (KST)"
  schedule_expression = "cron(0 2 ? * MON *)"
}

resource "aws_cloudwatch_event_target" "call_crawling_every_monday" {
  rule = aws_cloudwatch_event_rule.every_monday_11AM.name
  target_id = aws_lambda_function.crawling_lambda.id
  arn = aws_lambda_function.crawling_lambda.arn
}

resource "aws_lambda_permission" "allow_cloudwatch_to_call_lambda" {
  statement_id = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.crawling_lambda.function_name
  principal = "events.amazonaws.com"
  source_arn = aws_cloudwatch_event_rule.every_monday_11AM.arn
}