variable "service_name" {
  description = "ec2 service name"
}

variable "target_vpc" {
  description = "target vpc id"
}

variable "base_ami" {
  description = "base ami"
}

variable "instance_type" {
  description = "type"
  default = "t2.micro"
}

variable "public_subnets" {
  description = "subnet id"
}

variable "stack" {
  description = "tags stack"
}

variable "ebs_optimized" {}

variable "ec2_public_key" {
  
}

