.PHONY: gomodgen deploy delete build run clean

run: build
	dist/cli -text "$(TEXT)"

build:
	go build \
		-trimpath \
		-ldflags "-s -w -extldflags '-static'" \
		-o dist/cli \
		./cmd/cli && \
		chmod +x ./dist/cli

clean:
	rm -rf ./dist && rm cool_img.png

docker/run:
	sh docker_run.sh "$(TEXT)"

gomodgen:
	GO111MODULE=on go mod init

gcp/deploy:
	gcloud functions deploy printer --entry-point Printer --runtime go113 --trigger-http --max-instances 1 --allow-unauthenticated --memory 128MB --timeout 15

gcp/delete:
	gcloud functions delete printer
