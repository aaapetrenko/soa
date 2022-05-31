# Данные хранятся в SQLite базе данных todo.db в таблице todos
# CREATE TABLE "todos" "todo_id" INTEGER PRIMARY KEY, "description" TEXT NOT NULL, "status" TEXT NOT NULL );

import sqlite3
from flask import jsonify, url_for

DB_PATH = './todo.db'  # Update this path accordingly
NOTSTARTED = 'Not Started'


# Формирование структуры ответа на основе информации, представленной в базе данных
def make_public_todo(row):
    new_todo = {}
    for field in row.keys():
        # Предоставление URL объекта вместо его id - хорошая практика
        if field == 'todo_id':
            new_todo['uri'] = url_for('get_todo', todo_id=row['todo_id'], _external=True)
        else:
            new_todo[field] = row[field]

    return new_todo


def make_public_statistics(row):
    new_todo = {}
    for field in row.keys():
        if field == 'todo_id':
            new_todo['uri'] = url_for('get_statistics', todo_id=row['todo_id'], _external=True)
        else:
            new_todo[field] = row[field]

    return new_todo


def get_all_todos():
    try:
        conn = sqlite3.connect(DB_PATH)
        conn.row_factory = sqlite3.Row
        c = conn.cursor()
        c.execute('select * from todos')
        rows = c.fetchall()
        result = jsonify({'todos': list(map(make_public_todo, rows))})
        return result
    except Exception as e:
        print('Error: ', e)
        return None


def get_statistics(todo_id):
    try:
        conn = sqlite3.connect(DB_PATH)
        # Обеспечивает работу с названиями колонок в таблице
        conn.row_factory = sqlite3.Row
        c = conn.cursor()
        c.execute("select * from statistics where todo_id=?;", [todo_id])
        r = c.fetchone()
        return jsonify(make_public_statistics(r))
    except Exception as e:
        print('Error: ', e)
    return None


def get_all_statistics():
    try:
        conn = sqlite3.connect(DB_PATH)
        # Обеспечивает работу с названиями колонок в таблице
        conn.row_factory = sqlite3.Row
        c = conn.cursor()
        c.execute('select * from statistics')
        # Получаем список строк в перечислимом формате
        rows = c.fetchall()
        # С помощью функции map применяем функцию make_public_todo ко всем элементам rows
        result = jsonify({'statistics': list(map(make_public_statistics, rows))})
        return result
    except Exception as e:
        print('Error: ', e)
        return None


# Получить отдельный элемент
def get_todo(todo_id):
    try:
        conn = sqlite3.connect(DB_PATH)
        # Обеспечивает работу с названиями колонок в таблице
        conn.row_factory = sqlite3.Row
        c = conn.cursor()
        c.execute("select * from todos where todo_id=?;", [todo_id])
        r = c.fetchone()
        return jsonify(make_public_todo(r))
    except Exception as e:
        print('Error: ', e)
    return None


# Добавить элемент в таблицу
def add_to_list(description, name, gender, email, avatar):
    try:
        conn = sqlite3.connect(DB_PATH)
        c = conn.cursor()
        c.execute('insert into todos(description, name, gender, email, avatar, status) values(?,?,?,?,?,?)',
                  (description, name, gender, email, avatar, NOTSTARTED))
        conn.commit()
        result = get_todo(c.lastrowid)
        return result, c.lastrowid
    except Exception as e:
        print('Error: ', e)
        return None


def add_to_list_statistics(todo_id, num_wins, num_losses, num_sess):
    try:
        conn = sqlite3.connect(DB_PATH)
        c = conn.cursor()
        c.execute('insert or ignore into statistics(todo_id, numWins, numLosses, numSess) values(?,?,?,?)',
                  (todo_id, num_wins, num_losses, num_sess))
        print("aadda", c.lastrowid)
        conn.commit()
        result = get_statistics(c.lastrowid)
        return result
    except Exception as e:
        print('Error: ', e)
        return None


# Обновить элемент с todo_id в таблице
def update_todo(todo_id, description, status):
    try:
        conn = sqlite3.connect(DB_PATH)
        c = conn.cursor()
        c.execute('update todos set description=?, status=? where todo_id=?', (description, status, todo_id))
        conn.commit()
        result = get_todo(todo_id)
        return result
    except Exception as e:
        print('Error: ', e)
        return None


def update_todo_name(todo_id, name):
    try:
        conn = sqlite3.connect(DB_PATH)
        c = conn.cursor()
        c.execute('update todos set name=? where todo_id=?', (name, todo_id))
        conn.commit()
        result = get_todo(todo_id)
        return result
    except Exception as e:
        print('Error: ', e)
        return None


def update_todo_gender(todo_id, gender):
    try:
        conn = sqlite3.connect(DB_PATH)
        c = conn.cursor()
        c.execute('update todos set gender=? where todo_id=?', (gender, todo_id))
        conn.commit()
        result = get_todo(todo_id)
        return result
    except Exception as e:
        print('Error: ', e)
        return None


def update_todo_email(todo_id, email):
    try:
        conn = sqlite3.connect(DB_PATH)
        c = conn.cursor()
        c.execute('update todos set email=? where todo_id=?', (email, todo_id))
        conn.commit()
        result = get_todo(todo_id)
        return result
    except Exception as e:
        print('Error: ', e)
        return None


def update_todo_avatar(todo_id, avatar):
    try:
        conn = sqlite3.connect(DB_PATH)
        c = conn.cursor()
        c.execute('update todos set avatar=? where todo_id=?', (avatar, todo_id))
        conn.commit()
        result = get_todo(todo_id)
        return result
    except Exception as e:
        print('Error: ', e)
        return None


def update_statistics_numWins(todo_id, num_wins):
    try:
        conn = sqlite3.connect(DB_PATH)
        c = conn.cursor()
        c.execute('update statistics set numWins=? where todo_id=?', (num_wins, todo_id))
        conn.commit()
        result = get_statistics(todo_id)
        return result
    except Exception as e:
        print('Error: ', e)
        return None


def update_statistics_numLosses(todo_id, num_losses):
    try:
        conn = sqlite3.connect(DB_PATH)
        c = conn.cursor()
        c.execute('update statistics set numLosses=? where todo_id=?', (num_losses, todo_id))
        conn.commit()
        result = get_statistics(todo_id)
        return result
    except Exception as e:
        print('Error: ', e)
        return None


def update_statistics_numSess(todo_id, num_sess):
    try:
        conn = sqlite3.connect(DB_PATH)
        c = conn.cursor()
        c.execute('update statistics set numSess=? where todo_id=?', (num_sess, todo_id))
        conn.commit()
        result = get_statistics(todo_id)
        return result
    except Exception as e:
        print('Error: ', e)
        return None


# Удалить элемент из таблицы по индексу
def remove_todo(todo_id):
    try:
        conn = sqlite3.connect(DB_PATH)
        c = conn.cursor()
        c.execute('DELETE FROM todos WHERE todo_id=?', [todo_id])
        c.execute('DELETE FROM statistics WHERE todo_id=?', [todo_id])
        conn.commit()
        return jsonify({'result': True})
    except Exception as e:
        print('Error: ', e)
        return None
