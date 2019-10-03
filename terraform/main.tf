provider "google" {
  project     = "mass-sample-gitops-tf"
  region      = "asia-northeast1"
}

resource "google_compute_instance" "vm_instance" {
  name         = "terraform-instance"
  machine_type = "f1-micro"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network       = "default"
    access_config {
    }
  }
}
