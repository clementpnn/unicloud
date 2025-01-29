variable "project_id" {
  description = "ID du projet GCP"
  type        = string
  default     = "axial-reference-447820-m6"
}

variable "default_region" {
  type        = string
  description = "Région par défaut GCP"
  default     = "europe-west9"
}

variable "region_fr" {
  description = "Région GCP pour la France"
  type        = string
  default     = "europe-west9"
}

variable "region_de" {
  description = "Région GCP pour l'Allemagne"
  type        = string
  default     = "europe-west3"
}

variable "machine_type" {
  type    = string
  default = "e2-small"
}

variable "node_count" {
  type    = number
  default = 1
}

variable "ansible_user" {
  description = "Utilisateur pour la connexion SSH sur la VM"
  type        = string
  default     = "admin"
}

variable "adiarra_email" {
  description = "Email de Adiarra"
  type        = string
  sensitive   = true
}

variable "avisage_email" {
  description = "Email de Avisage"
  type        = string
  sensitive   = true
}

variable "mdesruets_email" {
  description = "Email de Mdesruets"
  type        = string
  sensitive   = true
}
