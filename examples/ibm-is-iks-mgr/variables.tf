variable "zone" {
  default = "us-east-1"
}

variable "ssh_public_key" {
  default = "~/.ssh/id_rsa.pub"
}

variable "image" {
  default = "ibm-ubuntu-18-04-1-minimal-amd64-2"
}

variable "profile" {
  default = "cx2-2x4"
}

variable "vpc" {
  default = "jtp1"
}

variable "name" {
}