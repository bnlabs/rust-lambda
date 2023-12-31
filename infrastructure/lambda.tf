resource "aws_lambda_function" "rust_lambda" {
  function_name = "rust-lambda"
  role          = aws_iam_role.rust_lambda_iam_role.arn
  handler       = "bootstrap"
  runtime       = "provided.al2"
  timeout       = 30
  source_code_hash = filebase64sha256("../bootstrap.zip")
  filename      = "../bootstrap.zip"
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

resource "aws_iam_role_policy" "lambda_logging" {
  name = "lambda_logging_policy"
  role = aws_iam_role.rust_lambda_iam_role.id

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        Resource = "arn:aws:logs:*:*:*"
      },
    ],
  })
}

resource "aws_iam_role_policy" "lambda_websocket_policy" {
  name = "lambda_websocket_policy"
  role = aws_iam_role.rust_lambda_iam_role.id

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Action = "execute-api:ManageConnections",
        Resource = "arn:aws:execute-api:us-east-1:915898657279:o6xs157o67/v1/POST/@connections/*"
      },
    ],
  })
}


