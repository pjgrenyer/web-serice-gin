output "app_url" {
  value = "https://${heroku_app.default.name}.herokuapp.com"
}

output "git_url" {
  value = heroku_app.default.git_url
}