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

# resource "aws_lambda_permission" "api_gw_lambda" {
#   statement_id  = "AllowExecutionFromAPIGateway"
#   action        = "lambda:InvokeFunction"
#   function_name = aws_lambda_function.rust_lambda.function_name
#   principal     = "apigateway.amazonaws.com"

#   # Note: The source ARN is for the method within your API Gateway that triggers the Lambda.
#   # Adjust the ARN as needed based on your API Gateway configuration.
#   source_arn = "${aws_api_gateway_rest_api.my_api.execution_arn}/*/*/*"
# }