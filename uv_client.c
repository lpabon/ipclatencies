#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <uv.h>

const char *socket_path = "go.sock";
int loop = 0;

void
on_close(uv_handle_t *handle) {
    free(handle);
}

void
on_read(uv_stream_t *handle,
        ssize_t nread,
        const uv_buf_t* buf) {


    if (nread < 0) {
        uv_close((uv_handle_t *)handle, on_close);
        return;
    }

    loop++;
    fprintf(stderr, "Loop: %d <%lld>", loop, nread);
    free(buf->base);
}

void
on_alloc(uv_handle_t *handle,
        size_t size,
        uv_buf_t *buf) {
    buf->base = malloc(4096);
    buf->len = 4096;
}

void
after_write(uv_write_t *req,
        int status) {
    free(req);
}

void
after_pipe_connect(uv_connect_t* req, int status) {

	int i, count;
    uv_write_t *write_req;
    uv_buf_t buf;
    char buffer[4096];

    if (status < 0) {
        abort();
    }

    printf("Connected\n");

	count = 10000;
    buf.base = buffer;
    buf.len = 4096;

	for (i = 0; i < count; i++) {

        write_req = malloc(sizeof(uv_write_t));

        printf("-->");

        uv_write(write_req,
                req->handle,
                &buf,
                1,
                after_write);

    }
    uv_read_start(req->handle, on_alloc, on_read);

    free(req);
}

int
main(int argc, char **argv) {


    uv_pipe_t q;
    uv_connect_t *connect_req;

    printf("Staring\n");

    uv_pipe_init(uv_default_loop(), &q, 0);
    
    connect_req = (uv_connect_t *)malloc(sizeof(uv_connect_t));

    uv_pipe_connect(connect_req,
            &q,
            socket_path,
            after_pipe_connect);

    uv_run(uv_default_loop(), UV_RUN_DEFAULT);
    uv_loop_close(uv_default_loop());

    printf("Leaving\n");

    return 0;
}
