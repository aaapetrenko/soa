import imghdr
from pathlib import Path

from fpdf import FPDF

import helper
from flask import Flask, abort, request, make_response
from datetime import datetime
import re
import fpdf

app = Flask(__name__)


@app.route("/")
def home():
    return "Hello, Flask!"


def add_image(pdf, image_path):
    pdf.image(image_path, x=10, y=8, w=100)
    return pdf


def add_text(pdf, header, text):
    pdf.cell(100, 5, txt=header + text, border=0, ln=1, align='', fill=False, link='')


@app.route('/statistics/api/v1.0/todos/<int:todo_id>', methods=['GET'])
def get_statistics(todo_id):
    response = helper.get_statistics(todo_id)
    if response is None:
        abort(404)
    response.headers.set('uri', '/statistics/api/v1.0/todos/' + str(todo_id))
    return response


@app.route('/statistics/api/v1.0/todos', methods=['GET'])
def get_all_statistics():
    response = helper.get_all_statistics()
    if response is None:
        abort(404)
    response.headers.set('uri', '/statistics/api/v1.0/todos')
    return response


@app.route('/todoapp/api/v1.0/todos/<int:todo_id>', methods=['GET'])
def get_todo(todo_id):
    response = helper.get_todo(todo_id)
    user_statistic = helper.get_statistics(todo_id).get_json()
    if response is None:
        abort(404)

    pdf = FPDF()
    pdf.add_page()
    if '.' in response.get_json()['avatar']:
        if imghdr.what(response.get_json()['avatar']) == 'jpeg' or imghdr.what(response.get_json()['avatar']) == 'png':
            pdf = add_image(pdf, response.get_json()['avatar'])
    pdf.set_font("Arial", size=12)
    add_text(pdf, "USER ID: ", str(todo_id))
    add_text(pdf, "USER description: ", response.get_json()['description'])
    add_text(pdf, "NAME: ", response.get_json()['name'])
    add_text(pdf, "GENDER: ", response.get_json()['gender'])
    add_text(pdf, "E-mail: ", response.get_json()['email'])
    add_text(pdf, "numWins: ", str(user_statistic['numWins']))
    add_text(pdf, "numLosses: ", str(user_statistic['numLosses']))
    add_text(pdf, "numSess: ", str(user_statistic['numSess']))
    pdf.output("simple_demo.pdf")
    response = make_response(pdf.output(dest='S').encode('latin-1'))
    response.headers.set('Content-Disposition', 'attachment', filename='name' + '.pdf')
    response.headers.set('Content-Type', 'application/pdf')
    response.headers.set('uri', '/todoapp/api/v1.0/todos/' + str(todo_id))
    return response


@app.route('/todoapp/api/v1.0/todos', methods=['GET'])
def get_all_todos():
    all_clients = helper.get_all_todos().get_json()
    all_statistics = helper.get_all_statistics().get_json()['statistics']
    pdf = FPDF()
    i = 0
    for el in all_clients['todos']:
        id_in_stat = el['uri'].split('/')[-1]
        pdf.add_page()
        if '.' in el['avatar']:
            if imghdr.what(el['avatar']) == 'jpeg' or imghdr.what(el['avatar']) == 'png':
                pdf = add_image(pdf, el['avatar'])
        pdf.set_font("Arial", size=12)

        add_text(pdf, "USER ID: ", str(id_in_stat))
        add_text(pdf, "USER description: ", el['description'])
        add_text(pdf, "NAME: ", el['name'])
        add_text(pdf, "GENDER: ", el['gender'])
        add_text(pdf, "E-mail: ", el['email'])
        if i < len(all_statistics):
            add_text(pdf, "numWins: ", str(all_statistics[i]['numWins']))
            add_text(pdf, "numLosses: ", str(all_statistics[i]['numLosses']))
            add_text(pdf, "numSess: ", str(all_statistics[i]['numSess']))
        print(el)
        i += 1
    response = make_response(pdf.output(dest='S').encode('latin-1'))
    response.headers.set('Content-Disposition', 'attachment', filename='name' + '.pdf')
    response.headers.set('Content-Type', 'application/pdf')
    response.headers.set('uri', '/todoapp/api/v1.0/todos')
    return response


def add_statistics(statistics_id):
    response = helper.add_to_list_statistics(statistics_id, 0, 0, 0)
    if response is None:
        abort(400)
    return response


@app.route('/todoapp/api/v1.0/todos', methods=['POST'])
def add_todo():
    if not request.json or not 'description' in request.json:
        abort(400)

    request_json = request.get_json()
    print(request_json)
    name = request_json.get('name', 'Name')
    gender = request_json.get('gender', '')
    email = request_json.get('email', '')
    avatar = request_json.get('avatar', '/Users/anastasiapetrenko/PycharmProjects/SeaBattle/RestHW/ittenVes170221.jpg')

    description = request.get_json()['description']
    response = helper.add_to_list(description, name, gender, email, avatar)
    if response[0] is None:
        abort(400)
    print(response[1])
    add_statistics(response[1])
    return response


@app.route('/statistics/api/v1.0/todos/<int:todo_id>', methods=['PUT'])
def update_statistics(todo_id):
    response = helper.get_statistics(todo_id).get_json()
    description_user = helper.get_todo(todo_id).get_json()['description']
    if description_user == 'reader':
        abort(400)
    print(response)
    if 'numWins' in request.json:
        response = helper.update_statistics_numWins(todo_id, request.get_json()['numWins'])
        if response is None:
            abort(400)
    if 'numLosses' in request.json:
        response = helper.update_statistics_numLosses(todo_id, request.get_json()['numLosses'])
        if response is None:
            abort(400)
    if 'numSess' in request.json:
        response = helper.update_statistics_numSess(todo_id, request.get_json()['numSess'])
        if response is None:
            abort(400)

    if response is None:
        abort(400)

    return response


@app.route('/todoapp/api/v1.0/todos/<int:todo_id>', methods=['PUT'])
def update_todo(todo_id):
    response = helper.get_todo(todo_id)
    if response is None:
        abort(404)

    if not request.json:
        abort(400)
    if not 'description' in request.json:
        abort(400)
    if response.get_json()['description'] == 'reader':
        abort(400)
    if not 'status' in request.json:
        abort(400)

    if 'name' in request.json:
        response = helper.update_todo_name(todo_id, request.get_json()['name'])
        if response is None:
            abort(400)
    if 'gender' in request.json:
        response = helper.update_todo_gender(todo_id, request.get_json()['gender'])
        if response is None:
            abort(400)
    if 'email' in request.json:
        response = helper.update_todo_email(todo_id, request.get_json()['email'])
        if response is None:
            abort(400)
    if 'avatar' in request.json:
        response = helper.update_todo_avatar(todo_id, request.get_json()['avatar'])
        if response is None:
            abort(400)

    if response is None:
        abort(400)

    return response


@app.route('/todoapp/api/v1.0/todos/<int:todo_id>', methods=['DELETE'])
def delete_task(todo_id):
    response = helper.get_todo(todo_id)
    if not response.get_json()['description']:
        abort(400)
    if response.get_json()['description'] == 'reader':
        abort(400)
    if response is None:
        abort(404)

    response = helper.remove_todo(todo_id)

    return response
