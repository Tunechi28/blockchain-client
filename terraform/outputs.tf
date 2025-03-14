output "ecs_service_names" {
  description = "List of ECS service names"
  value       = { for k, v in aws_ecs_service.main : k => v.name }
}

output "ecs_cluster_id" {
  description = "ECS Cluster ID"
  value       = aws_ecs_cluster.main.id
}

output "security_group_ids" {
  description = "Security group IDs for ECS services"
  value       = { for k, v in aws_security_group.ecs_service : k => v.id }
}

output "subnet_ids" {
  description = "Map of public subnet IDs"
  value       = { for k, v in aws_subnet.public : k => v.id }
}
