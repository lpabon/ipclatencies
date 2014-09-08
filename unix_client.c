#include <sys/socket.h>
#include <sys/un.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <stdint.h>
#include "tm.h"

char *socket_path = "go.sock";

int main(int argc, char *argv[]) {
  struct sockaddr_un addr;
  int fd,rc;
  char buf[4096];
  int len;
  int64_t d = 0;
  tm_ty ts, te;
  int i, count;

  if (argc > 1) socket_path=argv[1];

  if ( (fd = socket(AF_UNIX, SOCK_STREAM, 0)) == -1) {
    perror("socket error");
    exit(-1);
  }

  memset(&addr, 0, sizeof(addr));
  addr.sun_family = AF_UNIX;
  strncpy(addr.sun_path, socket_path, sizeof(addr.sun_path)-1);

  if (connect(fd, (struct sockaddr*)&addr, sizeof(addr)) == -1) {
    perror("connect error");
    exit(-1);
  }

  count = 100000;
  for (i=0; i<count; i++) {
      TM_NOW(ts);
      len = write(fd, buf, sizeof(buf));
      if (len < 1) {
          printf("unable to write\n");
          exit(0);
      }

      len = read(fd, buf, len);
      if (len < 1) {
        printf("unable to read\n");
        exit(0);
      }
      TM_NOW(te);

      d += TM_DURATION_NSEC(te, ts);

      usleep(10);
  }
  printf("Latency: %lld ns\n", d/((int64_t)count));


  return 0;
}