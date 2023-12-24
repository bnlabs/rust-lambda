resource "aws_s3_bucket" "terraform_state_bucket"{
    bucket = "rust-lambda-tf-state-bucket"

    lifecycle {
      prevent_destroy = true
    }
}

# resource "aws_s3_bucket_acl" "terraform_state_bucket_access_control_list" {
#   bucket = aws_s3_bucket.terraform_state_bucket.id
#   acl    = "private"
# }

resource "aws_s3_bucket_server_side_encryption_configuration" "terraform_state_bucket_serverside_encryption" {
  bucket = aws_s3_bucket.terraform_state_bucket.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm     = "AES256"
    }
  }
}

resource "aws_s3_bucket_versioning" "versioning" {
  bucket = aws_s3_bucket.terraform_state_bucket.id
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_dynamodb_table" "terraform_lock"{
    name = "terraform-state-locking"
    billing_mode = "PAY_PER_REQUEST"
    hash_key = "LockID"

    attribute {
      name = "LockID"
      type = "S"
    }
}