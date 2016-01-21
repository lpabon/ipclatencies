#include <sys/socket.h>
#include <sys/un.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <stdint.h>
#include <sys/uio.h>
#include "tm.h"

char *socket_path = "go.sock";

int main(int argc, char *argv[])
{
	struct sockaddr_un addr;
	int fd, rc;
	char buf[4096];
    char one[2048];
    char two[2048];
	int len;
	int64_t d = 0;
	int64_t wd = 0;
	int64_t rd = 0;
	tm_ty ts, te;
	tm_ty wte;
	int i, count;

	if (argc > 1)
		socket_path = argv[1];

	if ((fd = socket(AF_UNIX, SOCK_STREAM, 0)) == -1) {
		perror("socket error");
		exit(-1);
	}

	memset(&addr, 0, sizeof(addr));
	addr.sun_family = AF_UNIX;
	strncpy(addr.sun_path, socket_path, sizeof(addr.sun_path) - 1);

	if (connect(fd, (struct sockaddr *)&addr, sizeof(addr)) == -1) {
		perror("connect error");
		exit(-1);
	}

	count = 100000;
	for (i = 0; i < count; i++) {
        struct iovec iov[2];

        iov[0].iov_base = one;
        iov[0].iov_len = 2048;
        iov[1].iov_base = two;
        iov[1].iov_len = 2048;

		TM_NOW(ts);
		//len = write(fd, buf, sizeof(buf));
        len = writev(fd, iov, 2);
		if (len < 1) {
			printf("unable to write\n");
			exit(0);
		}
		TM_NOW(wte);
		wd += TM_DURATION_NSEC(wte, ts);

		len = read(fd, buf, sizeof(buf));
		if (len < 1) {
			printf("unable to read\n");
			exit(0);
		}
		TM_NOW(te);
		d += TM_DURATION_NSEC(te, ts);
		rd += TM_DURATION_NSEC(te, wte); 
	}
	printf("Latency: %lld ns\n", d / ((int64_t) count));
	printf("R Latency: %lld ns\n", rd / ((int64_t) count));
	printf("W Latency: %lld ns\n", wd / ((int64_t) count));



	return 0;
}
