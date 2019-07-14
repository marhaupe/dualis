build: 
	GOOS=darwin go build -o dualis_maco 
	GOOS=linux go build -o dualis_linux
	GOOS=windows go build -o dualis_windows

clean:
	rm -rf dist