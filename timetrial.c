#include <sys/socket.h>
#include <sys/un.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <stdint.h>
#include "tm.h"

char *socket_path = "go.sock";

int main(int argc, char *argv[]) {
	tm_ty ts, te;
  	int i, count;
  	int64_t d;

  	count = 10000000;
  	for (i=0; i<count; i++) {
  		TM_NOW(ts);
  		TM_NOW(te);
  		d += TM_DURATION_NSEC(te, ts);
  	}
  	printf("Time: %lld ns\n", d/((int64_t)count));

  	return 0;
}