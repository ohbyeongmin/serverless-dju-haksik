# Multi-account 세팅을 위해서 각 account마다 필요한 정책을 생성
# 공통 리소스 모듈

resource "aws_iam_policy" "policy" {
  name        = "assume-${var.aws_account}-${var.subject}-policy"
  path        = "/"
  description = "${var.aws_account}-${var.subject}-policy"
  policy      = data.aws_iam_policy_document.policy-document.json
}

data "aws_iam_policy_document" "policy-document" {
  statement {
    effect = "Allow"

    resources = var.resources

    actions = [
      "sts:AssumeRole",
    ]
  }
}