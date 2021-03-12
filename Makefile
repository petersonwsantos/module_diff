.PHONY: all clean

all:
	cd library; \
	GOOS=linux   GOARCH=amd64   go build -o diff     diff.go; \
	GOOS=linux   GOARCH=amd64   go build -o cca_diff_linux_x86_64       cca_diff.go; \
	GOOS=windows GOARCH=amd64   go build -o cca_diff_win32nt_64-bit.exe cca_diff.go
	
clean:
	rm -f library/cca_diff_*
