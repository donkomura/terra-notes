terraform {
  backend "gcs" {
    bucket = "donkomura-playground-tfstate"
    prefix = "terraform/state"
  }
}
