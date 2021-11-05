############## Dayone DevOps Group ##################
resource "aws_iam_group" "dh_devops" {
  name = "dh_devops"
}

resource "aws_iam_group_membership" "dh_devops" {
  name = aws_iam_group.dh_devops.name

  users = [
    aws_iam_user.admin_dh.name,
    aws_iam_user.dev_dh.name
  ]

  group = aws_iam_group.dh_devops.name
}

#######################################################

########### dh DevOps users #####################

resource "aws_iam_user" "admin_dh" {  
  name = "admin@dh.com"       # Edit this value to the username you want to use 
}

resource "aws_iam_user" "dev_dh" {
  name = "dev@dh.com"
}

#######################################################

############### DevOps admin Policy ##################

resource "aws_iam_user_policy" "dh_devops" {
  name  = "dh_devops"
  user = aws_iam_user.admin_dh.name

  policy = jsonencode({
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "*"
            ],
            "Resource": [
                "*"
            ]
        }
    ]
  })
}
######################################################


########### DevOps Assume Policies ####################
resource "aws_iam_group_policy_attachment" "dh_devops" {
  count      = length(var.userassume_policy_dh_devops)
  group      = aws_iam_group.dh_devops.name
  policy_arn = var.userassume_policy_dh_devops[count.index]
}

variable "userassume_policy_dh_devops" {
  description = "IAM Policy to be attached to user"
  type        = list(string)

  default = [
    # Please change <account_id> to the real account id number of id account 
    "arn:aws:iam::937681887038:policy/assume-dh-prod-admin-policy", # Add admin policy to black group user 
  ]
}

#######################################################


############### MFA Manager ###########################
resource "aws_iam_group_policy_attachment" "dh_devops_rotatekeys" {
  group      = aws_iam_group.dh_devops.name
  policy_arn = aws_iam_policy.RotateKeys.arn
}

resource "aws_iam_group_policy_attachment" "dh_devops_selfmanagemfa" {
  group      = aws_iam_group.dh_devops.name
  policy_arn = aws_iam_policy.SelfManageMFA.arn
}

resource "aws_iam_group_policy_attachment" "dh_devops_forcemfa" {
  group      = aws_iam_group.dh_devops.name
  policy_arn = aws_iam_policy.ForceMFA.arn
}

#######################################################