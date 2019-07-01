import socket

host = "127.0.0.1"
port = 12000

client = socket.socket(socket.AF_INET, socket.SOCK_STREAM) # Create Object

# Connecting to server
# This inst. does not need for UDP connection.
client.connect((host, port)) # Connecting to server

print("[Client] Connection established.")

client.send("Hi! This is client. Do you hear me, server?") # Send message

response = client.recv(4096) # Set receive as a power of 2 that you want

print("[Client] Received -> {0}".format(response))

