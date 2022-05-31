import socket
import threading

users = {}
room_users = {}


class Server:
    def __init__(self):
        self.ip = socket.gethostbyname(socket.gethostname())
        print('Running on IP: ' + self.ip)
        while 1:
            try:
                self.port = 8080  # int(input('Enter port number to run on --> '))

                self.s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
                self.s.bind((self.ip, self.port))

                break
            except:
                print("Couldn't bind to that port")

        self.connections = []
        self.accept_connections()

    def accept_connections(self):
        print('Running on port: ' + str(self.port))
        self.s.listen(100)

        while True:
            c, addr = self.s.accept()

            self.connections.append(c)

            threading.Thread(target=self.handle_client, args=(c, addr,)).start()
            print('new client', addr, c)
            users[c] = ''
            room_users[c] = ''

    def broadcast(self, sock, data):
        for client in self.connections:
            if client != self.s and client != sock:
                if client in room_users and sock in room_users:
                    if room_users[client] == room_users[sock]:
                        try:
                            client.send(data)
                        except:
                            pass
                elif client in room_users:
                    try:
                        client.send(data)
                    except:
                        pass

    def parser_name(self, data, c):
        data_dec = data
        if b'&&&' in data_dec:
            data_dec = data_dec.decode()
            if data_dec.split('&&&')[0] == 'NAME':
                users[c] = data_dec.split('&&&')[1].split('_')[0]
                room_users[c] = data_dec.split('&&&')[1].split('_')[1]
                print(users[c], room_users[c])
                return True
            return False
        return False

    def parser_list_of_user(self, data):
        data_dec = data
        if b'&&&' in data_dec:
            data_dec = data_dec.decode()
            if data_dec.split('&&&') == ['', 'LIST', '']:
                return True
            return False
        return False

    def parser_leave(self, c):
        users.pop(c)
        room_users.pop(c)

    def handle_client(self, c, addr):
        while 1:
            try:
                data = c.recv(1024)
                if b'&&&LEAVE&&&' in data:
                    self.parser_leave(c)
                    data_tmp = data.split(b'&&&')
                    if data_tmp[0] != b'':
                        self.broadcast(c, data_tmp[0])
                    if data_tmp[2] != b'':
                        self.broadcast(c, data_tmp[2])
                    continue
                if b'&&&LIST&&&' in data:
                    user_list = users.values()
                    user_list = '\n'.join(user_list)
                    c.send(user_list.encode())
                    data_tmp = data.split(b'&&&')
                    if data_tmp[0] != b'':
                        self.broadcast(c, data_tmp[0])
                    if data_tmp[2] != b'':
                        self.broadcast(c, data_tmp[2])
                    continue
                tmp = self.parser_name(data, c)
                if tmp:
                    print('New user')
                if tmp is False:
                    self.broadcast(c, data)

            except socket.error:
                c.close()


server = Server()
