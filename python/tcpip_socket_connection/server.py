import socket

host = "127.0.0.1"
port = 12000

serversock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
serversock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
serversock.bind((host, port)) # Binding with IP address and port
serversock.listen(10) # Listen port... (arg means the max number of queues)

print("[Server] Waiting for connections...")
clientsock, client_address = serversock.accept() # Store data if connection establised

while True:
    rcvmsg = clientsock.recv(1024)
    if rcvmsg != "":
        break

print("[Server] Received -> {0}".format(rcvmsg))
s_msg = "Hi! This is server. Thanks for sending message <3"

clientsock.sendall(s_msg) # resend message

clientsock.close()
