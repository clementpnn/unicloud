variable "project_id" {
  description = "ID du projet GCP"
  type        = string
  default     = "axial-reference-447820-m6"
}

variable "region" {
  description = "Région GCP"
  type        = string
  default     = "europe-west3"
}

variable "postgres_password" {
  description = "Mot de passe pour la base de données PostgreSQL"
  type        = string
}
