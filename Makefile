all:
	go build daemon.go
	go build client.go

install: all
	cp daemon $(DESTDIR)
	cp client $(DESTDIR)

clean:
	rm -f daemon client *.snap
