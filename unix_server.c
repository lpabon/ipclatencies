#include <stdio.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>

int server(int socket);

char *socket_path = "go.sock";

int main(int argc, char *argv[])
{
	struct sockaddr_un addr;
	int fd, client, pid;

	if (argc > 1) {
		socket_path = argv[1];
	}

	if ((fd = socket(AF_UNIX, SOCK_STREAM, 0)) == -1) {
		perror("socket error");
		exit(-1);
	}

	memset(&addr, 0, sizeof(addr));
	addr.sun_family = AF_UNIX;
	strncpy(addr.sun_path, socket_path, sizeof(addr.sun_path) - 1);

	unlink(socket_path);

	if (bind(fd, (struct sockaddr *)&addr, sizeof(addr)) == -1) {
		perror("bind error");
		exit(-1);
	}

	if (listen(fd, 10) == -1) {
		perror("listen error");
		exit(-1);
	}

	while (1) {
		if ((client = accept(fd, NULL, NULL)) == -1) {
			perror("accept error");
			continue;
		}

		pid = fork();
		if (0 == pid) {
			/* child */
			server(client);
			exit(0);
		} else if (pid < 0) {
			/* parent */
			close(client);
			close(fd);
			exit(-1);
		} else {
			close(client);
		}
	}

	return 0;
}

int server(int socket)
{
	char buf[4096];
	int len;

	while (1) {
		len = read(socket, buf, sizeof(buf));
		if (len < 0) {
			printf("Unable to read from socket\n");
			close(socket);
			return -1;
		} else if (len == 0) {
			printf("Close: EOF\n");
			close(socket);
			return 0;
		}

		len = write(socket, buf, len);
		if (len < 0) {
			printf("Unable to write to socket\n");
			close(socket);
			return -1;
		} else if (len == 0) {
			printf("Write close EOF\n");
			close(socket);
			return 0;
		}
	}
}
