resource "aws_lambda_function" "rust_lambda" {
  function_name = "rust-lambda"
  role          = aws_iam_role.rust_lambda_iam_role.arn
  handler       = "bootstrap"
  runtime       = "provided.al2"
  timeout       = 30

  filename      = "./target/lambda/rust-lambda/lambda.zip"
}

resource "aws_iam_role" "rust_lambda_iam_role" {
  name = "rust_lambda_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com"
        },
      },
    ],
  })
}