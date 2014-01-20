CFLAGS=-g -Wall
LDFLAGS=-lpulse

audio:audio.go libdde-pulse.so
	$(MAKE) libdde-pulse.so
	go build
libdde-pulse.so:dde-pulse.c dde-pulse.h
	$(CC) -fPIC -shared -o $@ $^ $(LDFLAGS) $(CFLAGS)
clean:
	go clean
	rm *.so *.o
