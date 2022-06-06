resource "heroku_app" "default" {
    name = var.app_name
    region = "eu"

    config_vars = {
    }
}
