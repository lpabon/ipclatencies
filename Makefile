GCC = gcc
all: cus cuc ctt

hwclient: Makefile hwclient.c tm.h
	$(GCC) -O2 -o hwclient hwclient.c -lzmq

hwserver: Makefile hwserver.c tm.h
	$(GCC) -O2 -o hwserver hwserver.c -lzmq

cus: Makefile unix_server.c tm.h
	$(GCC) -O2 -o cus unix_server.c

cuc: Makefile unix_client.c tm.h
	$(GCC) -O2 -o cuc unix_client.c

ctt: Makefile tm.h timetrial.c
	$(GCC) -O2 -o ctt timetrial.c

clean:
	rm -f hwclient hwserver cus cuc

.PHONY: clean
