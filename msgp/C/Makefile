GCC = gcc

all: msp_cuc

msp_cuc: Makefile unix_client_msp.c tm.h
	$(GCC) -O2 -o msp_cuc unix_client_msp.c -lmsgpack

clean:
	rm -f msp_cuc

.PHONY: clean
