GO=/opt/local/bin/go

all: merge-ParseRequestURI

clean:
	rm -f merge-ParseRequestURI
world:
	make clean
	make all
	
merge-ParseRequestURI: merge-ParseRequestURI.go
	$(GO) build merge-ParseRequestURI.go
