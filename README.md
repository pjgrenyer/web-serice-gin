docker build -t web-serice-gin .

docker run -it --rm --name web-serice-gin -p 8080:8080 web-serice-gin
