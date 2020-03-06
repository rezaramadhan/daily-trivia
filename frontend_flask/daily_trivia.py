from flask import Flask, render_template

app = Flask(__name__)

@app.route('/')
def home():
    return render_template('index.html')

@app.route('/week')
@app.route('/month')
@app.route('/year')
def coming_soon():
    return render_template('coming_soon.html')
