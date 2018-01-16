FROM hashicorp/terraform:light
ADD release/linux/amd64/ /usr/local/terraform-plugins/
