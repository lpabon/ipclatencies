#include <sys/socket.h>
#include <sys/un.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <stdint.h>
#include "tm.h"
#include <msgpack.h>

char *socket_path = "/tmp/go.sock";

typedef struct {
    uint32_t id;
    uint64_t block;
    uint64_t nblocks;
    char *path;
} pkt_t;

void
pkt_free(pkt_t *p) {
    free(p->path);
}

msgpack_sbuffer *msgpack_tutorial(void)
{
	/* msgpack::sbuffer is a simple buffer implementation. */
	msgpack_sbuffer *sbuf;
	sbuf = msgpack_sbuffer_new();

	/* serialize values into the buffer using msgpack_sbuffer_write callback function. */
	msgpack_packer pk;
	msgpack_packer_init(&pk, sbuf, msgpack_sbuffer_write);

#if 0
	/*
	 * {
	 *   "Id" : uint32,
	 *   "Block" : uint64,
	 *   "NBlocks" : uint64,
	 *   "Path" : string
	 * }
	 */
	msgpack_pack_map(&pk, 4);

	/* id */
	msgpack_pack_str(&pk, 2);
	msgpack_pack_str_body(&pk, "Id", 2);
	msgpack_pack_uint32(&pk, 32);

	/* Block */
	msgpack_pack_str(&pk, 5);
	msgpack_pack_str_body(&pk, "Block", 5);
	msgpack_pack_uint64(&pk, 123456);

	/* NBlocks */
	msgpack_pack_str(&pk, 7);
	msgpack_pack_str_body(&pk, "NBlocks", 7);
	msgpack_pack_uint64(&pk, 789);

	/* Path */
	msgpack_pack_str(&pk, 4);
	msgpack_pack_str_body(&pk, "Path", 4);
	msgpack_pack_str(&pk, strlen("/usr/home/here"));
	msgpack_pack_str_body(&pk, "/usr/home/here", strlen("/usr/home/here"));
#endif

    msgpack_pack_array(&pk, 4);
	msgpack_pack_uint32(&pk, 0);
	msgpack_pack_uint64(&pk, 123456);
	msgpack_pack_uint64(&pk, 789);
	msgpack_pack_str(&pk, strlen("/usr/home/here"));
	msgpack_pack_str_body(&pk, "/usr/home/here", strlen("/usr/home/here"));

    /*
	msgpack_pack_array(&pk, 3);
	msgpack_pack_int(&pk, 1);
	msgpack_pack_true(&pk);
	msgpack_pack_str(&pk, 7);
	msgpack_pack_str_body(&pk, "example", 7);
    */

	/* deserialize the buffer into msgpack_object instance. */
	/* deserialized object is valid during the msgpack_zone instance alive. */
	msgpack_zone mempool;
	msgpack_zone_init(&mempool, 2048);

	msgpack_object deserialized;
	msgpack_unpack(sbuf->data, sbuf->size, NULL, &mempool, &deserialized);

	/* print the deserialized object. */
	msgpack_object_print(stdout, deserialized);
	puts("");

	msgpack_zone_destroy(&mempool);

	return sbuf;
}

msgpack_sbuffer *
marshal(pkt_t* p) {

	/* msgpack::sbuffer is a simple buffer implementation. */
	msgpack_sbuffer *sbuf;
	sbuf = msgpack_sbuffer_new();

	/* serialize values into the buffer using msgpack_sbuffer_write callback function. */
	msgpack_packer pk;
	msgpack_packer_init(&pk, sbuf, msgpack_sbuffer_write);

    msgpack_pack_array(&pk, 4);

	msgpack_pack_uint32(&pk, p->id);
	msgpack_pack_uint64(&pk, p->block);
	msgpack_pack_uint64(&pk, p->nblocks);

	msgpack_pack_str(&pk, strlen(p->path));
	msgpack_pack_str_body(&pk, p->path, strlen(p->path));

    return sbuf;
}

void
unmarshal(pkt_t* p, char *buf, size_t len) {
    int i;
	msgpack_object deserialized;
	msgpack_zone mempool;

	msgpack_zone_init(&mempool, 2048);

	msgpack_unpack(buf, len, NULL, &mempool, &deserialized);
    /*
    msgpack_object_print(stdout, deserialized);
    puts("");
    */

    if (deserialized.type == MSGPACK_OBJECT_ARRAY) {

        for (i = 0; i < deserialized.via.array.size; i++) {
            switch(i) {
                case 0:
                    p->id = deserialized.via.array.ptr[i].via.u64;
                    break;

                case 1:
                    p->block = deserialized.via.array.ptr[i].via.u64;
                    break;

                case 2:
                    p->nblocks = deserialized.via.array.ptr[i].via.u64;
                    break;

                case 3:
                    if (p->path != NULL) {
                        free(p->path);
                    }
                    p->path = strndup(deserialized.via.array.ptr[i].via.str.ptr, 
                            deserialized.via.array.ptr[i].via.str.size);
                    break;
            }
        }
    }
	msgpack_zone_destroy(&mempool);
}

int main(int argc, char *argv[])
{
	struct sockaddr_un addr;
	int fd, rc;
	char buf[4096];
	int len;
	int64_t d = 0;
	int64_t wd = 0;
	int64_t rd = 0;
	int64_t md = 0;
	tm_ty ts, te;
	tm_ty wte;
	tm_ty ms, me;
	int i, count;

	msgpack_sbuffer *sbuf;

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
	// Create a pack
	//sbuf = msgpack_tutorial();
    pkt_t p;
    p.id = 0;
    p.block = 0;
    p.nblocks = 1;
    p.path = strdup("/hello");
    

	count = 1000000;
	for (i = 0; i < count; i++) {

        /*
         * Send
         */
        sbuf = marshal(&p);
		TM_NOW(ts);
		len = write(fd, sbuf->data, sbuf->size);
		if (len < 1) {
			printf("unable to write\n");
			exit(0);
		}
		TM_NOW(wte);
        msgpack_sbuffer_destroy(sbuf);
		wd += TM_DURATION_NSEC(wte, ts);

        /*
         * Recv
         */
		len = read(fd, buf, sizeof(buf));
		if (len < 1) {
			printf("unable to read\n");
			exit(0);
		}
		TM_NOW(te);
		d += TM_DURATION_NSEC(te, ts);
		rd += TM_DURATION_NSEC(te, wte);

		TM_NOW(ms);
        unmarshal(&p, buf, len);
		TM_NOW(me);

		md += TM_DURATION_NSEC(ms, me);
	}
	printf("Latency: %lld ns\n", d / ((int64_t) count));
	printf("R Latency: %lld ns\n", rd / ((int64_t) count));
	printf("W Latency: %lld ns\n", wd / ((int64_t) count));
	printf("MPCK: %lld ns\n", md / ((int64_t) count));
    pkt_free(&p);


	return 0;
}
