docker.build:
	docker build -t extractor:v0.0.1 .

docker.run:
	docker run -it -p 8080:8080 extractor:v0.0.1 

test.success_with_data:
	k6 run test/success_with_data.js

test.success_with_warning:
	k6 run test/success_with_warning.js

test.error_not_found:
	k6 run test/error_not_found.js