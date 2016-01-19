#include <sys/socket.h>
#include <sys/un.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <stdint.h>
#include <pthread.h>
#include "tm.h"

#define MESSAGES 10000
char *socket_path = "go.sock";

void *reader(void *arg)
{

	int64_t d = 0;
	tm_ty ts, te;
	tm_ty *send_ts;
	int i;
	int fd = *(int *)arg;
	char buf[4096];
	int len;

	send_ts = (tm_ty *) buf;
	for (i = 0; i < MESSAGES; i++) {
		len = read(fd, buf, sizeof(buf));
		if (len < 1) {
			printf("unable to read\n");
			exit(0);
		}
		TM_NOW(te);

		ts.tv_nsec = send_ts->tv_nsec;
		ts.tv_sec = send_ts->tv_sec;
		d += TM_DURATION_NSEC(te, ts);
	}

	printf("Latency: %lld ns\n", d / ((int64_t) MESSAGES));
	return NULL;
}

int main(int argc, char *argv[])
{
	struct sockaddr_un addr;
	int fd, rc;
	char buf[4096];
	int len;
	tm_ty *send_ts;
	pthread_t tid;
	int i;

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

	/*
	 * start server 
	 */
	pthread_create(&tid, NULL, reader, (void *)&fd);

	send_ts = (tm_ty *) buf;
	for (i = 0; i < MESSAGES; i++) {
		clock_gettime(CLOCK_MONOTONIC, send_ts);
		len = write(fd, buf, sizeof(buf));
		if (len < 1) {
			printf("unable to write\n");
			exit(0);
		}
        usleep(100);
	}

    printf("Done\n");

	pthread_join(tid, NULL);

	return 0;
}
