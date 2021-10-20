output "parameters" {
  description = "The parameters that were created"
  value = { for key, param in var.parameters : key => aws_ssm_parameter.parameters[key].value }
  sensitive = true
}
