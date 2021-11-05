variable "aws_region" {
  description = "The AWS region to deploy the shard storage layer into"
}

variable "account_id" {
  description = "The AWS account number"
}

variable "assume_role_arn" {
  description = "The role to assume when accessing the AWS API."
  default     = ""
}

variable "atlantis_user" {
  description = "The username that will be triggering atlantis commands. This will be used to name the session when assuming a role. More information - https://github.com/runatlantis/atlantis#assume-role-session-names"
  default     = "atlantis_user"
}

variable "remote_state_region" {
  default = "ap-northeast-2"
}

variable "remote_state_bucket" {
  default = ""
}


variable "prod_account_id" {
  description = "The AWS account number for produdction"
}

variable "remote_state_key_map" {
  type = map(string)

  default = {
    "kms_apne2" = "art/terraform/kms/art-id/id_apnortheast2/terraform.tfstate"
  }
}