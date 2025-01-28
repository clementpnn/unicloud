resource "google_container_cluster" "de" {
  name                     = "gke-de"
  location                 = var.region_de
  project                  = var.project_id
  remove_default_node_pool = true
  initial_node_count       = 1
}

resource "google_container_node_pool" "de_pool" {
  name     = "de-pool"
  project  = var.project_id
  location = var.region_de
  cluster  = google_container_cluster.de.name

  node_count = var.node_count

  node_config {
    machine_type = var.machine_type
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
}
