import socket
import sys
import os

server_address = 'go.sock'
# Make sure the socket does not already exist
try:
    os.unlink(server_address)
except OSError:
    if os.path.exists(server_address):
        raise

# Create a UDS socket
sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)

# Bind the socket to the port
print >>sys.stderr, 'starting up on %s' % server_address
sock.bind(server_address)

# Listen for incoming connections
sock.listen(10)

while True:
    # Wait for a connection
    connection, client_address = sock.accept()
    try:
        # Receive the data in small chunks and retransmit it
        while True:
            data = connection.recv(4096)
            if data:
                connection.sendall(data)
            else:
                break

    finally:
        # Clean up the connection
        connection.close()
