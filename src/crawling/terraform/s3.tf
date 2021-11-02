resource "random_pet" "dh_bucket_name" {
  prefix = "dh-bot-bucket"
  length = 2
}

resource "aws_s3_bucket" "dh_bucket" {
  bucket = random_pet.dh_bucket_name.id

  tags = {
    Name = "dh_s3_bucket"
  }
}

