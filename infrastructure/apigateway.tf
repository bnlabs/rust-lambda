# resource "aws_api_gateway_rest_api" "my_api" {
#   name        = "my api gateway"
#   description = "Managed by Terraform"
# }

# resource "aws_api_gateway_resource" "my_resource" {
#   rest_api_id = aws_api_gateway_rest_api.my_api.id
#   parent_id   = aws_api_gateway_rest_api.my_api.root_resource_id
#   path_part   = "chat"
# }

# resource "aws_api_gateway_method" "my_method" {
#   rest_api_id   = aws_api_gateway_rest_api.my_api.id
#   resource_id   = aws_api_gateway_resource.my_resource.id
#   http_method   = "GET" # or POST, PUT, DELETE...
#   authorization = "NONE"
# }

# resource "aws_api_gateway_integration" "lambda_integration" {
#   rest_api_id = aws_api_gateway_rest_api.my_api.id
#   resource_id = aws_api_gateway_resource.my_resource.id
#   http_method = aws_api_gateway_method.my_method.http_method
#   integration_http_method = "POST"
#   type        = "AWS_PROXY"
#   uri         = aws_lambda_function.rust_lambda.invoke_arn
# }

# resource "aws_api_gateway_deployment" "my_deployment" {
#   depends_on = [
#     aws_api_gateway_integration.lambda_integration
#   ]
#   rest_api_id = aws_api_gateway_rest_api.my_api.id
#   stage_name  = "v1"
# }

# output "invoke_url" {
#   value = "${aws_api_gateway_deployment.my_deployment.invoke_url}/${aws_api_gateway_resource.my_resource.path}"
# }

# // enable cors
# resource "aws_api_gateway_method" "options_method" {
#   rest_api_id   = aws_api_gateway_rest_api.my_api.id
#   resource_id   = aws_api_gateway_resource.my_resource.id
#   http_method   = "OPTIONS"
#   authorization = "NONE"
# }

# resource "aws_api_gateway_method_response" "options_method_response" {
#   rest_api_id = aws_api_gateway_rest_api.my_api.id
#   resource_id = aws_api_gateway_resource.my_resource.id
#   http_method = aws_api_gateway_method.options_method.http_method
#   status_code = "200"
#   response_models = {
#     "application/json" = "Empty"
#   }

#   response_parameters = {
#     "method.response.header.Access-Control-Allow-Headers" = true
#     "method.response.header.Access-Control-Allow-Methods" = true
#     "method.response.header.Access-Control-Allow-Origin"  = true
#   }
# }


# Define the WebSocket API
resource "aws_apigatewayv2_api" "websocket_api" {
  name          = "my-websocket-api"
  protocol_type = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
}

# Define the connect route
resource "aws_apigatewayv2_route" "connect_route" {
  api_id    = aws_apigatewayv2_api.websocket_api.id
  route_key = "$connect"
  target    = "integrations/${aws_apigatewayv2_integration.websocket_lambda_integration.id}"
}

# Define the disconnect route
resource "aws_apigatewayv2_route" "disconnect_route" {
  api_id    = aws_apigatewayv2_api.websocket_api.id
  route_key = "$disconnect"
  target    = "integrations/${aws_apigatewayv2_integration.websocket_lambda_integration.id}"
}

resource "aws_apigatewayv2_route" "default_route" {
  api_id    = aws_apigatewayv2_api.websocket_api.id
  route_key = "$default"
  target    = "integrations/${aws_apigatewayv2_integration.websocket_lambda_integration.id}"
}

resource "aws_apigatewayv2_integration" "websocket_lambda_integration" {
  api_id             = aws_apigatewayv2_api.websocket_api.id
  integration_type   = "AWS_PROXY"
  integration_uri    = aws_lambda_function.rust_lambda.invoke_arn
  description        = "Lambda integration"
  payload_format_version = "1.0"
}

# Link the integrations to the routes
resource "aws_apigatewayv2_integration_response" "integration_response" {
  api_id    = aws_apigatewayv2_api.websocket_api.id
  integration_id = aws_apigatewayv2_integration.websocket_lambda_integration.id
  integration_response_key = "$default"
}

# Deployment of the WebSocket API
resource "aws_apigatewayv2_deployment" "websocket_deployment" {
  api_id      = aws_apigatewayv2_api.websocket_api.id
  description = "WebSocket Deployment"

  depends_on = [
    aws_apigatewayv2_route.connect_route,
    aws_apigatewayv2_route.disconnect_route,
    aws_apigatewayv2_route.default_route,
    aws_apigatewayv2_integration.websocket_lambda_integration
  ]
}

# Stage for deployment
resource "aws_apigatewayv2_stage" "websocket_stage" {
  api_id      = aws_apigatewayv2_api.websocket_api.id
  name        = "v1"
  # deployment_id = aws_apigatewayv2_deployment.websocket_deployment.id
  auto_deploy = true
}

# Output the WebSocket URL
output "websocket_url" {
  value = aws_apigatewayv2_api.websocket_api.api_endpoint
}

# Lambda permission to allow WebSocket API to invoke it
resource "aws_lambda_permission" "websocket_lambda_permission" {
  statement_id  = "AllowExecutionFromWebSocketAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.rust_lambda.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.websocket_api.execution_arn}/*/*"
}
