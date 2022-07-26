.PHONY: build
build:
	DOCKER_BUILDKIT=1
	docker build --target bin --output bin/ .

.PHONY: clean
clean:
	rm bin/bb