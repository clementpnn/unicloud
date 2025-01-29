resource "kubernetes_namespace" "prod_de" {
  provider = kubernetes.de
  metadata {
    name = "prod-de"
  }
}

resource "kubernetes_config_map" "db_init_sql_de" {
  provider = kubernetes.de

  metadata {
    name      = "db-init-de"
    namespace = kubernetes_namespace.prod_de.metadata[0].name
  }

  data = {
    "init.sql" = file("${path.module}/../sql/init.sql")
  }
}

resource "kubernetes_deployment" "front_de" {
  provider = kubernetes.de
  metadata {
    name      = "frontend"
    namespace = kubernetes_namespace.prod_fr.metadata[0].name
    labels = {
      app = "frontend"
    }
  }

  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "frontend"
      }
    }
    template {
      metadata {
        labels = {
          app = "frontend"
        }
      }
      spec {
        container {
          name  = "frontend"
          image = "my-registry.com/my-frontend:latest"
          port {
            container_port = 80
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "front_de_svc" {
  provider = kubernetes.de
  metadata {
    name      = "frontend-service"
    namespace = kubernetes_namespace.prod_fr.metadata[0].name
  }
  spec {
    selector = {
      app = "frontend"
    }
    port {
      name        = "http"
      port        = 80
      target_port = 80
    }
    type = "LoadBalancer"
  }
}

resource "kubernetes_deployment" "back_de" {
  provider = kubernetes.de
  metadata {
    name      = "backend"
    namespace = kubernetes_namespace.prod_fr.metadata[0].name
    labels = {
      app = "backend"
    }
  }

  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "backend"
      }
    }
    template {
      metadata {
        labels = {
          app = "backend"
        }
      }
      spec {
        container {
          name  = "backend"
          image = "my-registry.com/my-backend:latest"
          port {
            container_port = 8080
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "back_de_svc" {
  provider = kubernetes.de
  metadata {
    name      = "backend-service"
    namespace = kubernetes_namespace.prod_fr.metadata[0].name
  }
  spec {
    selector = {
      app = "backend"
    }
    port {
      name        = "http"
      port        = 8080
      target_port = 8080
    }
    type = "ClusterIP"
  }
}

resource "kubernetes_deployment" "db_de" {
  provider = kubernetes.de
  metadata {
    name      = "database"
    namespace = kubernetes_namespace.prod_fr.metadata[0].name
    labels = {
      app = "database"
    }
  }

  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "database"
      }
    }
    template {
      metadata {
        labels = {
          app = "database"
        }
      }
      spec {
        container {
          name  = "db"
          image = "mysql:8.0"
          env {
            name  = "MYSQL_ROOT_PASSWORD"
            value = "root"
          }
          port {
            container_port = 3306
          }
          volume_mount {
            name       = "db-init-volume"
            mount_path = "/docker-entrypoint-initdb.d"
            read_only  = true
          }
        }
        volume {
          name = "db-init-volume"

          config_map {
            name = kubernetes_config_map.db_init_sql.metadata[0].name
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "db_de_svc" {
  provider = kubernetes.de
  metadata {
    name      = "database-service"
    namespace = kubernetes_namespace.prod_fr.metadata[0].name
  }
  spec {
    selector = {
      app = "database"
    }
    port {
      name        = "db"
      port        = 3306
      target_port = 3306
    }
    type = "ClusterIP"
  }
}
