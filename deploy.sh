heroku container:login
docker build --no-cache -t web-service-gin:latest .
heroku container:push web --app pg-web-service-gin
sleep 10
heroku container:release web --app pg-web-service-gin
heroku open --app pg-web-service-gin
heroku logs --tail --app pg-web-service-gin
