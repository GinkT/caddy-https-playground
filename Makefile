build:
	docker build -t cc-test .

run:
	docker run -p 12321:12321 cc-test