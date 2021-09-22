variable "prefix" {
  description = "The prefix to use for all SSM Parameters"
  type        = string
}

variable "github_repository" {
  description = "The Github repository to add secrets to for running automation"
  type        = string
  default     = "ssm2env"
}

variable "type" {
  description = "The type of SSM Parameter to create"
  type        = string
  default     = "String"
}

variable "tags" {
  description = "Tags to attach to the SSM Parameters"
  type        = map(string)
  default     = {}
}

variable "parameters" {
  description = "The parameters to create for testing purposes"
  type        = map(string)
  default = {
    foo = "<secure-string>"
    bar = "-"
    baz = "&escaped"
  }
}
