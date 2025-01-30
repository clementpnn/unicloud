output "database_connection_name" {
  description = "Connection string for the Cloud SQL instance"
  value       = google_sql_database_instance.postgres_instance.connection_name
}

output "database_public_ip" {
  description = "Public IP of the Cloud SQL instance"
  value       = google_sql_database_instance.postgres_instance.public_ip_address
}
