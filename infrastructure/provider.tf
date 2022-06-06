provider "heroku" {
  email   = var.heroku_email
  api_key = var.heroku_api_key
}

terraform {
  required_providers {
    heroku = {
      source  = "heroku/heroku"
      version = "~> 4.0"
    }
  }

  backend "remote" {
    organization = "Paul_Grenyer"

    workspaces {
      name = "web-service-gin"
    }
  }
}
