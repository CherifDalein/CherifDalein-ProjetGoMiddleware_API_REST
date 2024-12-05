resource "openstack_compute_instance_v2" "basic" {
  name            = "basic"
  image_id        = data.openstack_images_image_v2.deb.id
  flavor_name       = "v1.m1024.d10"
  security_groups = ["default"]
  config_drive = true
  user_data       = file("init.yaml")

  metadata = {
    this = "that"
  }

  network {
    name = "public"
  }
}


data "openstack_networking_secgroup_v2" "secgroup_1" {
  name        = "default"
}

resource "openstack_networking_secgroup_rule_v2" "secgroup_rule_1" {
  direction         = "ingress"
  ethertype         = "IPv6"
  protocol          = "tcp"
  port_range_min    = 22
  port_range_max    = 22
  remote_ip_prefix  = "::/0"
  security_group_id = data.openstack_networking_secgroup_v2.secgroup_1.id
}
