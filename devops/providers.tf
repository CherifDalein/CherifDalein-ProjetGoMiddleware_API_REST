# On détermine la version du provider OpenStack à utiliser
terraform {
  required_providers {
    openstack = {
      source = "terraform-provider-openstack/openstack"
    }
  }
  required_version = ">= 1.0.0"
}

# On demande à OpenTofu d'utiliser le provider téléchargé à l'instant
provider "openstack" {
  cloud = "openstack"
}
