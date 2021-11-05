#
# dh-prod administrator
#
resource "aws_iam_role" "assume_dh_prod_admin" {
  name = "assume-dh-prod-admin"
  path = "/"
  max_session_duration = "43200"
  assume_role_policy = jsonencode({
  "Version": "2012-10-17",
  "Statement": [
            {
            "Sid": "",
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::${var.id_account_id}:root"
            },
            "Action": "sts:AssumeRole"
            }
        ]
    })
}

resource "aws_iam_role_policy" "assume_dh_prod_admin" {
  name = "assume-dh-prod-admin-passrole"
  role = aws_iam_role.assume_dh_prod_admin.id

  policy = jsonencode({
  "Statement": [
            {
            "Sid": "AllowIAMPassRole",
            "Action": [
                "iam:PassRole"
            ],
            "Effect": "Allow",
            "Resource": "*"
            }
        ]
    })
}

resource "aws_iam_role_policy_attachment" "assume_dh_prod_admin" {
  role       = aws_iam_role.assume_dh_prod_admin.id
  policy_arn = "arn:aws:iam::aws:policy/AdministratorAccess"
}