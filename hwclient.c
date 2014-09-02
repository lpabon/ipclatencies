//  Hello World client
#include <zmq.h>
#include <string.h>
#include <stdio.h>
#include <stdint.h>
#include <unistd.h>
#include "tm.h"

int main (void)
{
    unsigned char buffer[8];
    int i, count;
    int64_t d = 0;
    tm_ty ts, te;

    printf ("Connecting to hello world server...\n");
    void *context = zmq_ctx_new ();
    void *requester = zmq_socket (context, ZMQ_REQ);
    zmq_connect (requester, "ipc://latency.ipc");
    

    count = 10000;
    for (i=0; i<count; i++) {
        TM_NOW(ts);
        zmq_send (requester, buffer, sizeof(buffer), 0);
        zmq_recv (requester, buffer, sizeof(buffer), 0);
        TM_NOW(te);

        d += TM_DURATION_NSEC(te, ts);
    }

    printf("Latency: %lld ns\n", d/((int64_t)count));

    zmq_close (requester);
    zmq_ctx_destroy (context);
    return 0;
}
