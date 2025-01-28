resource "google_container_cluster" "fr" {
  name                     = "gke-fr"
  location                 = var.region_fr
  project                  = var.project_id
  remove_default_node_pool = true
  initial_node_count       = 1

}

resource "google_container_node_pool" "fr_pool" {
  name     = "fr-pool"
  project  = var.project_id
  location = var.region_fr
  cluster  = google_container_cluster.fr.name

  node_count = var.node_count

  node_config {
    machine_type = var.machine_type
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
}
