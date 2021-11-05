variable "aws_account" {
  description = "aws account name"
}

variable "subject" {
  description = ""
}

variable "resources" {
  type        = list
  description = "List of affected resources."
  default     = ["*"]
}