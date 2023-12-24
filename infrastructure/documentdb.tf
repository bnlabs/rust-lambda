# resource "aws_docdb_cluster" "Rust_Lambda_Db" {
#   cluster_identifier      = "my-docdb-cluster"
#   engine                  = "docdb"
#   master_username         = "place_holder_username"
#   master_password         = "place_holder_password"
#   backup_retention_period = 5
#   preferred_backup_window = "07:00-09:00"
#   skip_final_snapshot     = true
# }