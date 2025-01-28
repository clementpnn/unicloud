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
