import socket

c = socket.socket()
c.connect(("localhost", 8080))

while True:    
    # Send msg
    msg = input("Client -> ")
    c.send(msg.encode(encoding="utf-8"))
    
    # Recieve msg
    server_msg = c.recv(1024).decode()
    print("Server <- ", server_msg)
c.close()