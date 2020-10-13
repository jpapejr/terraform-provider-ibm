data "ibm_is_vpc" "vpc" {
  name = var.vpc
}

data "ibm_is_subnet" "subnet" {
  name            = var.subnet
}

data "ibm_is_ssh_key" "sshkey" {
  name       = "jtpape"
}

data "ibm_is_security_group" "sg"{
  name      = var.sg
}

resource "ibm_is_instance" "instance1" {
  name    = "instance1"
  image   = var.image
  profile = var.profile

  primary_network_interface {
    subnet = data.ibm_is_subnet.subnet.id
  }

  vpc       = data.ibm_is_vpc.vpc.id
  zone      = var.zone
  keys      = [data.ibm_is_ssh_key.sshkey.id]
}

resource "ibm_is_security_group_network_interface_attachment" "sgnic1" {
  depends_on        = [ ibm_is_instance.instance1 ]
  security_group    = data.ibm_is_security_group.sg.id
  network_interface = ibm_is_instance.instance1.primary_network_interface[0].id
}

resource "ibm_is_floating_ip" "floatingip1" {
  name   = "${var.name}-ip"
  target = ibm_is_instance.instance1.primary_network_interface[0].id
}

