import pyaudio
import wave
import socket

filename = "recorded.wav"
chunk = 1024
FORMAT = pyaudio.paInt16
channels = 1
sample_rate = 40960
record_seconds = 5



class Client:
    def __init__(self):
        self.s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        while 1:
            try:
                self.target_ip = input('Enter IP address of server --> ')
                self.target_port = int(input('Enter target port of server --> '))
                self.s.connect((self.target_ip, self.target_port))
                self._user_name = input('Enter user name --> ')
                self._room = input('Enter room number --> ')
                print(':) Hello,', self._user_name)
                tmp = self._user_name.encode()
                self.s.send(b'NAME&&&' + tmp + b'_' + self._room.encode())
                break
            except:
                print(":) Couldn't connect to server")

    def start_recoding(self):
        p = pyaudio.PyAudio()
        stream = p.open(format=FORMAT,
                        channels=channels,
                        rate=sample_rate,
                        input=True,
                        # input_device_index=2,
                        output=True,
                        frames_per_buffer=chunk)
        frames = []
        print(":) Recording...")
        for i in range(int(40960 / chunk * record_seconds)):
            data = stream.read(chunk)
            # stream.write(data)
            frames.append(data)
        print(":) Finished recording.")
        stream.stop_stream()
        stream.close()
        p.terminate()
        wf = wave.open(filename, "wb")
        wf.setnchannels(channels)
        wf.setsampwidth(p.get_sample_size(FORMAT))
        wf.setframerate(sample_rate)
        wf.writeframes(b"".join(frames))
        wf.close()
        self.send_data_to_serer()

    def send_data_to_serer(self):
        p = pyaudio.PyAudio()
        wf = wave.open('recorded.wav', 'rb')
        stream = p.open(format=p.get_format_from_width(wf.getsampwidth()),
                        channels=wf.getnchannels(),
                        rate=wf.getframerate(),
                        output=True)

        data = wf.readframes(chunk)
        gl_data = data
        while data != b'':
            data = wf.readframes(chunk)
            gl_data += data
            self.s.send(data)
        self.s.send('###END'.encode())
        stream.stop_stream()
        stream.close()
        p.terminate()
        pass

    def rcv_from_server(self):
        p = pyaudio.PyAudio()
        wf = wave.open('recorded.wav', 'rb')
        stream = p.open(format=p.get_format_from_width(wf.getsampwidth()),
                        channels=wf.getnchannels(),
                        rate=wf.getframerate(),
                        output=True)
        while 1:
            self.s.settimeout(100)
            data = self.s.recv(1024)
            if not data:
                break
            if b'###' in data:
                components = data.split(b'###')
                if components[1] == b'END':
                    stream.write(components[0])
                    break
            stream.write(data)

        stream.stop_stream()
        stream.close()
        p.terminate()

    def list_of_user(self):
        self.s.send(b'&&&LIST&&&')
        data = self.s.recv(1024)
        data = data.decode()
        print(':) List of users:\n', data)

    def destroy(self):
        self.s.send(b'&&&LEAVE&&&')
        self.s.close()


client = Client()
print('1) to get a list of users enter the command "users"')
print('2) to exit enter the command "exit"')
print('3) to receive a voice message enter the command "recv"')
print('4) to send a voice message enter the command "send"')
while True:
    a = input('Enter the command --> ')
    if a == 'exit':
        client.destroy()
        break
    elif a == 'send':
        client.start_recoding()
    elif a == 'recv':
        client.rcv_from_server()
    elif a == 'users':
        client.list_of_user()
