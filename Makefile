GCC = gcc
all: hwclient hwserver

hwclient: Makefile hwclient.c tm.h
	$(GCC) -O2 -o hwclient hwclient.c -lzmq 

hwserver: Makefile hwserver.c tm.h
	$(GCC) -O2 -o hwserver hwserver.c -lzmq 
