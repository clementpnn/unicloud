resource "google_sql_database_instance" "postgres_instance" {
  name             = "cloudsql-postgres"
  database_version = "POSTGRES_15"
  region           = var.region

  depends_on = [google_service_networking_connection.private_vpc_connection]

  settings {
    tier              = "db-custom-1-3840"
    availability_type = "REGIONAL"

    backup_configuration {
      enabled                        = true
      point_in_time_recovery_enabled = true
      location                       = var.region
    }

    ip_configuration {
      ipv4_enabled    = true
      private_network = google_compute_network.vpc.self_link
      authorized_networks {
        name  = "public-access"
        value = "0.0.0.0/0"
      }
    }
  }


}

resource "google_sql_database" "postgres_db" {
  name     = "mydatabase"
  instance = google_sql_database_instance.postgres_instance.name

  depends_on = [google_sql_database_instance.postgres_instance]
}

resource "google_sql_user" "postgres_user" {
  name     = "myuser"
  instance = google_sql_database_instance.postgres_instance.name
  password = var.postgres_password

  depends_on = [google_sql_database_instance.postgres_instance]
}

resource "null_resource" "import_sql" {
  depends_on = [google_sql_database.postgres_db]

  provisioner "local-exec" {
    command = <<EOT
      PGPASSWORD=${var.postgres_password} psql \
        -h $(gcloud sql instances describe cloudsql-postgres --format="value(ipAddresses[0].ipAddress)") \
        -U myuser \
        -d mydatabase \
        -f ../sql/init.sql
    EOT
  }
}
