//  Hello World server

#include <zmq.h>
#include <stdio.h>
#include <unistd.h>
#include <string.h>
#include <assert.h>
#include <stdint.h>

typedef struct {
    int32_t id, count, result;
} Message_t;

int main (void)
{
    char buffer [256];
    Message_t *msg = (Message_t *)buffer;

    //  Socket to talk to clients
    void *context = zmq_ctx_new ();
    void *responder = zmq_socket (context, ZMQ_REP);
    int rc = zmq_bind (responder, "ipc://latency.ipc");
    assert (rc == 0);

    while (1) {
        int len;

        len = zmq_recv (responder, buffer, sizeof(buffer), 0);
        msg->result = msg->count + 10;
        zmq_send (responder, buffer, len,  0);
    }
    return 0;
}
