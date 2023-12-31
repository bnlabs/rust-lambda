# resource "aws_dynamodb_table" "connectionId-dynamodb-table"{
#     name = "websocket-connectionIds"
#     billing_mode = "PAY_PER_REQUEST"
#     hash_key = "ConnectionID"

#     attribute {
#       name = "ConnectionID"
#       type = "S"
#     }
# }