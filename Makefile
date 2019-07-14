build:
	go build

release: 
	goreleaser

clean:
	rm -rf dist