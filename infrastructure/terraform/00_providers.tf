terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.0"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.0"
    }
  }
  required_version = ">= 1.3.0"
}

provider "google" {
  project = var.project_id
  region  = var.default_region
}

provider "kubernetes" {
}

provider "kubernetes" {
  alias                  = "fr"
  host                   = google_container_cluster.fr.endpoint
  token                  = data.google_client_config.default.access_token
  cluster_ca_certificate = base64decode(google_container_cluster.fr.master_auth.0.cluster_ca_certificate)
}

provider "kubernetes" {
  alias                  = "de"
  host                   = google_container_cluster.de.endpoint
  token                  = data.google_client_config.default.access_token
  cluster_ca_certificate = base64decode(google_container_cluster.de.master_auth.0.cluster_ca_certificate)
}

data "google_client_config" "default" {}
