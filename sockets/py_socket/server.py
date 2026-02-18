import socket

s = socket.socket()
# s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)

s.bind(("localhost", 8080))

print("Server is listening...!")
s.listen(5)

while True:
    client_socket, addr = s.accept()
    # print("client socket:", client_socket)
    print("Client connected:", addr)
    
    while True:
        # Recieve msg
        client_msg = client_socket.recv(1028).decode()
        print("Client <- ", client_msg)
        
        # Send msg
        server_msg = input("Server -> ")
        client_socket.send(server_msg.encode(encoding="utf-8"))
    
    client_socket.close()