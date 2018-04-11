# terraform-provider-rancheros

This terraform plugin can be used to access a Rancher API's instance. While this project violates a lot of the norm of terraform, it is created to assist in a zero touch turn up of a Rancher cluster. When using terraform to generate a HA Rancher cluster, localAuth must be set via API to prevent public access and an API access and secret key are required in order to use the offical rancher provider in terraform.

If you're like me and are using AWS to generate auto_scaling_groups and dynamically add workers, you'll need to create an API key in order to use the offical a Terraform plugin. While creation of API keys is possible with bash, Terraform does not currently support a way to import vars. See the workflow for more details
