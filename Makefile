GCC = gcc
all: mcus mcuc cus cuc ctt

hwclient: Makefile hwclient.c tm.h
	$(GCC) -O2 -o hwclient hwclient.c -lzmq

hwserver: Makefile hwserver.c tm.h
	$(GCC) -O2 -o hwserver hwserver.c -lzmq

mcus: Makefile unix_server_mt.c tm.h
	$(GCC) -O2 -pthread -o mcus unix_server_mt.c

mcuc: Makefile unix_client_mt.c tm.h
	$(GCC) -O2 -pthread -o mcuc unix_client_mt.c

cus: Makefile unix_server.c tm.h
	$(GCC) -O2 -o cus unix_server.c

cuc: Makefile unix_client.c tm.h
	$(GCC) -O2 -o cuc unix_client.c

ctt: Makefile tm.h timetrial.c
	$(GCC) -O2 -o ctt timetrial.c

clean:
	rm -f mcus mcuc hwclient hwserver cus cuc

.PHONY: clean
