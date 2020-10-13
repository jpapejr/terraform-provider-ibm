data "ibm_is_vpc" "vpc" {
  name = var.vpc
}

data "ibm_is_subnet" "subnet" {
  name            = var.subnet
  vpc             = ibm_is_vpc.vpc.id
}

data "ibm_is_ssh_key" "sshkey" {
  name       = "jtpape"
}

resource "ibm_is_instance" "instance1" {
  name    = "instance1"
  image   = var.image
  profile = var.profile

  primary_network_interface {
    subnet = ibm_is_subnet.subnet.id
  }

  vpc       = ibm_is_vpc.vpc.id
  zone      = var.zone
  keys      = [ibm_is_ssh_key.sshkey.id]
}
