# SSM2Env Integration Tests Infrastructure

This terraform code will create ssm parameters for running the integration tests.

## Prerequisites

* Terraform

## First Time Setup

1. Copy `terraform.tfvars.example` as `terraform.tfvars`.
1. Copy `versions.tf.example` as `versions.tf`. It is recommended to fill out a backend.
1. Initialize the terraform module: `terraform init`
1. Create a plan: `terraform plan -out=test-params.plan`
1. Apply the plan: `terraform apply test-params.plan`
1. Set the `TEST_SSM2ENV_INTEGRATION` environment variable to 0 to pull the parameters and allow the integration tests to be run
1. Refresh the environment variables: `direnv reload`
