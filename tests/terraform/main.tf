resource "aws_ssm_parameter" "parameters" {
  for_each = var.parameters
  name     = join("/", [var.prefix, each.key])
  type     = var.type
  value    = each.value
  tags     = var.tags
}

resource "github_actions_secret" "tfvars_secrets" {
  for_each = {
    prefix     = var.prefix
    type       = var.type
    tags       = jsonencode(var.tags)
    parameters = jsonencode(var.parameters)
  }

  repository      = var.github_repository
  secret_name     = "TF_VAR_${each.key}"
  plaintext_value = each.value
}

resource "github_actions_secret" "versions_secret" {
  repository      = var.github_repository
  secret_name     = "VERSIONS_TF_FILE"
  plaintext_value = base64encode(file("${path.module}/versions.tf"))
}
