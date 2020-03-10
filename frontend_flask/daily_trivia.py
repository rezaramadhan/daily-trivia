from flask import Flask, render_template

import requests
import json
import os

app = Flask(__name__)

API_HOST = os.environ.get("API_HOST", default='localhost')
API_PORT = os.environ.get("API_PORT", default='8080')

def get_trivia():
    url = f'http://{API_HOST}:{API_PORT}/today'
    resp = requests.get(url)
    json_content = json.loads(resp.content)
    return json_content

@app.route('/')
def home():
    data = get_trivia()
    return render_template('index.html', trivia=data['trivia'], d=data['day'], m=data['month'])

@app.route('/week')
@app.route('/month')
@app.route('/year')
def coming_soon():
    return render_template('coming_soon.html')
