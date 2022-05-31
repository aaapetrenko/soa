import os
import subprocess
import asyncio

import aiohttp

p1 = 'http://127.0.0.1:5000/todoapp/api/v1.0/todos'
p2 = 'http://127.0.0.1:5000/statistics/api/v1.0/todos'

d = {
    "name": "N",
    'gender': "f",
    "email": "a@mail.ru",
    "avatar": "/Users/anastasiapetrenko/PycharmProjects/SeaBattle/RestHW/ittenVes170221.jpg",
    "description": "editor"
}

a = {
    'numWins': 1,
    'numLosses': 1,
    'numSess': 1
}
async def get(uri):
    async with aiohttp.ClientSession() as session:
        async with session.get(uri) as response:
            print("Status:", response.status)
            print("Content-type:", response.headers['content-type'])
            if 'uri' not in response.headers:
                print('Error')
            else:
                print("uri", 'http://127.0.0.1:5000' + response.headers['uri'])


async def delete(uri):
    async with aiohttp.ClientSession() as session:
        async with session.delete(uri) as response:
            print("Status:", response.status)
            print("Content-type:", response.headers['content-type'])


async def post(uri, data):
    async with aiohttp.ClientSession() as session:
        async with session.post(uri, json=data) as response:
            print("Status:", response.status)
            print("Content-type:", response.headers['content-type'])


async def put(uri, data):
    async with aiohttp.ClientSession() as session:
        async with session.put(uri, json=data) as response:
            print("Status:", response.status)
            print("Content-type:", response.headers['content-type'])


loop = asyncio.get_event_loop()
i = 0
while True:
    command = input("Enter the command --> ")
    i += 1
    if i >= 50:
        break
    path = ''
    if command == 'Exit':
        break
    if 'Get' in command:
        if command == 'GetAllStat':
            path = 'http://127.0.0.1:5000/statistics/api/v1.0/todos'
        elif command == 'GetStat':
            number = input("Enter user id --> ")
            path = 'http://127.0.0.1:5000/statistics/api/v1.0/todos/' + str(number)
        elif command == 'GetAll':
            path = 'http://127.0.0.1:5000/todoapp/api/v1.0/todos'
        elif command == 'GetUser':
            number = input("Enter user id --> ")
            path = 'http://127.0.0.1:5000/todoapp/api/v1.0/todos/' + str(number)
        else:
            continue
        loop.run_until_complete(get(path))
    elif command == 'Delete':
        number = input("Enter user id --> ")
        path = 'http://127.0.0.1:5000/todoapp/api/v1.0/todos/' + str(number)
        loop.run_until_complete(delete(path))
    elif command == 'Post':
        path = 'http://127.0.0.1:5000/todoapp/api/v1.0/todos'
        loop.run_until_complete(post(path, d))
    elif command == 'PutStat':
        number = input("Enter user id --> ")
        path = 'http://127.0.0.1:5000/statistics/api/v1.0/todos/' + str(number)
        loop.run_until_complete(put(path, a))
    elif command == 'PutUserInfo':
        number = input("Enter user id --> ")
        path = 'http://127.0.0.1:5000/todoapp/api/v1.0/todos/' + str(number)
        loop.run_until_complete(put(path, d))

