LIBUVLIBS=$(shell pkg-config --libs libuv)
GCC = gcc
all: mcus mcuc cus cuc cuc_iov

mcus: Makefile unix_server_mt.c tm.h
	$(GCC) -O2 -pthread -o mcus unix_server_mt.c

mcuc: Makefile unix_client_mt.c tm.h
	$(GCC) -O2 -pthread -o mcuc unix_client_mt.c

cus: Makefile unix_server.c tm.h
	$(GCC) -O2 -o cus unix_server.c

cuc_iov: Makefile unix_client_iov.c tm.h
	$(GCC) -O2 -o cuc_iov unix_client.c

cuc: Makefile unix_client.c tm.h
	$(GCC) -O2 -o cuc unix_client.c

uvuc: Makefile uv_client.c tm.h
	$(GCC) -O0 -g -o uvuc uv_client.c $(LIBUVLIBS) 


clean:
	rm -f uvuc mcus mcuc cus cuc

.PHONY: clean
