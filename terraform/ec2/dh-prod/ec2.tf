resource "aws_security_group" "ec2" {
  name        = "${var.service_name}"
  description = "${var.service_name} Instance Security Group"
  vpc_id      = var.target_vpc

  ingress {
    from_port = 22
    to_port   = 22
    protocol  = "tcp"

    cidr_blocks = ["0.0.0.0/0"]

    description = "SSH port"
  }

  ingress {
    from_port = 8080
    to_port   = 8080
    protocol  = "tcp"

    cidr_blocks = ["0.0.0.0/0"]

    description = "jenkins port"
  }


  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Internal outbound traffic"
  }
}

resource "aws_key_pair" "dh-ec2-key" {
  key_name   = "dh-ec2-key"
  public_key = var.ec2_public_key
}

resource "aws_instance" "prod-ec2" {
  ami                    = var.base_ami
  instance_type          = var.instance_type
  subnet_id              = var.public_subnets[0]
  vpc_security_group_ids = [aws_security_group.ec2.id]
  ebs_optimized          = var.ebs_optimized
  key_name = aws_key_pair.dh-ec2-key.key_name

  iam_instance_profile = aws_iam_instance_profile.ec2-profile.name

  root_block_device {
    volume_size = "10"
  }

  tags = {
    Name  = "${var.service_name}-${var.stack}"
    app   = "${var.service_name}"
    stack = var.stack
  }

}

resource "aws_iam_role" "ec2-role" {
  name = "ec2-role"

  assume_role_policy = jsonencode({
  "Version": "2012-10-17",
  "Statement": [
      {
        "Action": "sts:AssumeRole",
        "Principal": {
          "Service": "ec2.amazonaws.com"
        },
        "Effect": "Allow",
        "Sid": ""
      }
    ]
  })
}

resource "aws_iam_instance_profile" "ec2-profile" {
  name = "ec2-profile"
  role = "${aws_iam_role.ec2-role.name}"
}

resource "aws_iam_role_policy" "ec2-policy" {
  name = "ec-2policy"
  role = "${aws_iam_role.ec2-role.id}"

  policy =  jsonencode({
  "Version": "2012-10-17",
  "Statement": [
      {
        "Action": [
          "ec2:AmazonEC2FullAccess"
        ],
        "Effect": "Allow",
        "Resource": "*"
      }
    ]
  })
}
