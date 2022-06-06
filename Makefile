PORT=8080

clean:
	rm -fdr dist

run:
	go run .

build: clean
	go build -v -o dist/app

docker:
	docker build -t web-serice-gin .

docker-run:
	docker run -it --rm --name web-serice-gin -p 8080:8080 web-serice-gin

publish: build
	heroku container:login
	docker build --no-cache -t web-service-gin:latest .
	heroku container:push web --app pg-web-service-gin

deploy: publish
	sleep 10
	heroku container:release web --app pg-web-service-gin
#	 heroku open --app pg-web-service-gin
	heroku logs --tail --app pg-web-service-gin

.DEFAULT_GOAL := run