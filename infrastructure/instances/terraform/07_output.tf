output "frontend_fr_lb_ip" {
  description = "IP publique du frontend FR"
  value       = kubernetes_service.front_fr_svc.status[0].load_balancer_ingress[0].ip
}

output "frontend_de_lb_ip" {
  description = "IP publique du frontend DE"
  value       = kubernetes_service.front_de_svc.status[0].load_balancer_ingress[0].ip
}
