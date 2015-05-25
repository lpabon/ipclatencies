#include <stdio.h>
#include <stdbool.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <pthread.h>
#include <sys/queue.h>

char *socket_path = "go.sock";

#ifdef LUIS
#define T(fmt, ...) do {\
    printf("%s:%d:%s(): " fmt "\n", __FILE__, __LINE__, __func__, ##__VA_ARGS__);\
    fflush(stdout);\
}while(0)
#else
#define T(fmt, ...) do { }while(0)
#endif

typedef struct {
    int socket;
    char buf[4096];
    int len;
} msg_t;


typedef struct {
    pthread_cond_t send_c;
    pthread_cond_t recv_c;
    pthread_mutex_t m;
    msg_t *msg;
} channel_t;

typedef struct {
    channel_t workers;
    channel_t reply;
} conn_t;

typedef struct {
    conn_t *conn;
    int socket;
} server_t;


void
ch_init(channel_t *ch) {
    ch->msg = NULL;
    pthread_mutex_init(&ch->m, NULL);
    pthread_cond_init(&ch->send_c, NULL);
    pthread_cond_init(&ch->recv_c, NULL);
}

void
ch_send(channel_t *ch, msg_t *msg) {

    pthread_mutex_lock(&ch->m);
    while (NULL != ch->msg) {
        pthread_cond_wait(&ch->send_c, &ch->m);
    }

    ch->msg = msg;

    pthread_cond_signal(&ch->recv_c);
    pthread_mutex_unlock(&ch->m);
}

msg_t *
ch_recv(channel_t *ch) {

    msg_t *msg;

    pthread_mutex_lock(&ch->m);
    while (NULL == ch->msg) {
        pthread_cond_wait(&ch->recv_c, &ch->m);
    }

    msg = ch->msg;
    ch->msg = NULL;

    pthread_cond_signal(&ch->send_c);
    pthread_mutex_unlock(&ch->m);

    return msg;
}

void *
worker(void *arg) {
    conn_t *conn = (conn_t *)arg;
    msg_t *msg;

    while (1) {
        /* get */
        msg = ch_recv(&conn->workers); 
        T("s:%d got message", msg->socket);

        /* Work */
        usleep(100);

        /* Send */
        T("s:%d @@send", msg->socket);
        ch_send(&conn->reply, msg);
        T("s:%d sent++", msg->socket);
    }

    return NULL;
}

void *
writer(void *arg) {
    conn_t *conn = (conn_t *)arg;
    msg_t *msg;
    int len;

    while (1) {
        msg = ch_recv(&conn->reply);
        T("msg ready to be sent");

        len = write(msg->socket, msg->buf, msg->len);
        if (len < 0) {
            printf("Unable to write to socket\n");
            close(msg->socket);
            return NULL;
        }
        else if (len == 0) {
            printf("Write close EOF\n");
            close(msg->socket);
            return NULL;
        }

        free(msg);
    }

    return NULL;
}

void *
reader(void *arg) {
    int i;
    msg_t *msg;
    server_t *s;
    conn_t *conn;

    /* setup conn */
    s = (server_t *)arg;
    conn = s->conn;

    T("start");
    while(1) {
        msg = (msg_t *)malloc(sizeof(msg_t));

        msg->socket = s->socket;
        msg->len = read(s->socket, msg->buf, sizeof(msg->buf));
        if (msg->len < 0) {
            printf("Unable to read from socket\n");
            close(s->socket);
            goto out;
        }
        else if (msg->len == 0) {
            printf("Close: EOF\n");
            close(s->socket);
            goto out;
        }

        T("s:%d send", s->socket);
        /* send */
        ch_send(&conn->workers, msg); 
        T("s:%d back", s->socket);
    }

out:
    free(arg);
    return NULL;
}

int main(int argc, char *argv[]) {
  struct sockaddr_un addr;
  int fd, client, pid;
    pthread_t tid;
    conn_t conn;
    int i;
    server_t *s;

  if (argc > 1) {
    socket_path=argv[1];
  }
  
    ch_init(&conn.workers);
    ch_init(&conn.reply);
    /* spawn */
    pthread_create(&tid, NULL, writer, &conn);

    // spawn workers
    for (i=0; i<32; i++) {
        pthread_create(&tid, NULL, worker, &conn);
    }

  if ( (fd = socket(AF_UNIX, SOCK_STREAM, 0)) == -1) {
    perror("socket error");
    exit(-1);
  }

  memset(&addr, 0, sizeof(addr));
  addr.sun_family = AF_UNIX;
  strncpy(addr.sun_path, socket_path, sizeof(addr.sun_path)-1);

  unlink(socket_path);

  if (bind(fd, (struct sockaddr*)&addr, sizeof(addr)) == -1) {
    perror("bind error");
    exit(-1);
  }

  if (listen(fd, 10) == -1) {
    perror("listen error");
    exit(-1);
  }

    while (1) {
        if ( (client = accept(fd, NULL, NULL)) == -1) {
          perror("accept error");
          continue;
        }

        s = (server_t *)malloc(sizeof (server_t));
        s->conn = &conn;
        s->socket = client;

        pthread_create(&tid, NULL, reader, s);

      }

  return 0;
}
