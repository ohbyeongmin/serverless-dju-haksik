output "backend_s3_name" {
    description = "Name of the S3 bucket used to store tfstate"

    value = random_pet.tfstate_bucket_name.id
}

