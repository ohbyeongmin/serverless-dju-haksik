resource "aws_s3_bucket" "s3" {
  bucket = "dh-${var.service_name}-${var.env_name}-bucket"

  tags = {
    Name = "dh_s3_${var.env_name}"
  }
}
