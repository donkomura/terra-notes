terraform {
  backend "gcs" {
    bucket = "donkomura-playground-tfstate"
    prefix = "env/dev"
  }
}
