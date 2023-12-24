terraform {
    backend "s3" {
        bucket  = "rust-lambda-tf-state-bucket"
        key     = "state/terraform.tfstate"
        region  = "us-east-1"
        encrypt = true
        dynamodb_table = "terraform-state-locking"
    }
}