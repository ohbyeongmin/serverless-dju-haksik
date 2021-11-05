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

  root_block_device {
    volume_size = "10"
  }

  tags = {
    Name  = "${var.service_name}-${var.stack}"
    app   = "${var.service_name}"
    stack = var.stack
  }

}