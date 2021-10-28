docker-build:
	docker build -t extractor:v0.0.1 .

docker-run:
	docker run -it -p 8080:8080 extractor:v0.0.1 
