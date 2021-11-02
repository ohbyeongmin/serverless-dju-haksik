output "logs_policy_arn" {
    description = "Arn of the logs policy"

    value = aws_iam_policy.logs_policy.arn
}