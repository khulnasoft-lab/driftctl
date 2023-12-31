provider "google" {}

terraform {
  required_version = "~> 0.15.0"
  required_providers {
    google = {
      version = "3.90.1"
    }
  }
}

resource "google_compute_address" "ip_address" {
  name   = "my-address"
  region = "us-central1"
}
