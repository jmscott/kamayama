GO=/opt/local/bin/go

all: upsert-ParseRequestURI

clean:
	rm -f upsert-ParseRequestURI
world:
	make clean
	make all
	
upsert-ParseRequestURI: upsert-ParseRequestURI.go
	$(GO) build upsert-ParseRequestURI.go
