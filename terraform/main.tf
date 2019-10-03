provider "google" {
  credentials = "${file("#{service_account_json_file}")}"
  project     = "mass-sample-gitops"
  region      = "asia-northeast1"
}

resource "google_compute_network" "sample-gitops" {
  name = "sample-gitops"
}
resource "google_compute_subnetwork" "development" {
  name          = "development"
  ip_cidr_range = "10.30.0.0/16"
  network       = "${google_compute_network.sample-gitops.name}"
  description   = "development"
  region        = "asia-northeast1"
}

resource "google_compute_firewall" "development" {
  name    = "development"
  network = "${google_compute_network.sample-gitops.name}"

  allow {
    protocol = "icmp"
  }

  allow {
    protocol = "tcp"
    ports    = ["22", "80", "443"]
  }

  target_tags = ["${google_compute_instance.development.tags}"]
}

resource "google_compute_instance" "development" {
  name         = "development"
  machine_type = "n1-standard-1"
  zone         = "asia-northeast1-c"
  description  = "sample-gitops"
  tags         = ["development", "mass"]

  disk {
    image = "ubuntu-os-cloud/ubuntu-1404-lts"
  }

  disk {
    type        = "local-ssd"
    scratch     = true
    auto_delete = true
  }

  network_interface {
    access_config {
    }

    subnetwork = "${google_compute_subnetwork.development.name}"
  }

  service_account {
    scopes = ["userinfo-email", "compute-ro", "storage-ro", "bigquery", "monitoring"]
  }

  scheduling {
    on_host_maintenance = "MIGRATE"
    automatic_restart   = true
  }
}
