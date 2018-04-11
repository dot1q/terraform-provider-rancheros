# terraform-provider-rancheros

## What's the point?
This terraform plugin can be used to access a Rancher API's instance. While this project violates a lot of the norm of terraform, it is created to assist in a zero touch turn up of a Rancher cluster. When using terraform to generate a HA Rancher cluster, localAuth must be set via API to prevent public access and an API access and secret key are required in order to use the offical rancher provider in terraform.

If you're like me and are using AWS to generate auto_scaling_groups and dynamically add workers, you'll need to create an API key in order to use the offical a Terraform plugin. While creation of API keys is possible with bash, Terraform does not currently support a way to import vars. See the workflow for more details

## Usage
* rancheros_apikey - CRUD for API keys
* rancheros_registrationtoken - Not implemented yet
* rancheros_localauthconfig - Not implemented yet

### rancheros_apikey
#### Example usage
```
resource "rancheros_apikey" "temp" {
  host = "https://rancherdev.aws.example.com:8443"
  name = "testApiNew"
  description = "CREATED BY TERRAFORM YO"
  access_key = "27FF9F387BF786AC755E"
  secret_key = "y7i5uXz5oLwzrFHooQiwkuDKLbvdatm5P75d7tLk"
}
```
#### Argument Reference
* host (required) - protocol://address:port of the rancher cluster or instance
* access_key (optional) - If auth is already enabled, you must specify a valid API access key to use
* secret_key (optional) - If auth is already enabled, you must specify a valid API secret key to use
* name (required) - The name associated with the API key
* description (optional) - The description field for rancher

#### Attribute Reference
* gen_access_key - The access key generated
* gen_secret_key - The secret key generated
