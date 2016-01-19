GCC = gcc
all: mcus mcuc cus cuc

mcus: Makefile unix_server_mt.c tm.h
	$(GCC) -O2 -pthread -o mcus unix_server_mt.c

mcuc: Makefile unix_client_mt.c tm.h
	$(GCC) -O2 -pthread -o mcuc unix_client_mt.c

cus: Makefile unix_server.c tm.h
	$(GCC) -O2 -o cus unix_server.c

cuc: Makefile unix_client.c tm.h
	$(GCC) -O2 -o cuc unix_client.c

clean:
	rm -f mcus mcuc cus cuc

.PHONY: clean
