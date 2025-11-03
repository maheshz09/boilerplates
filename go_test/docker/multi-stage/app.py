#this is test commit 
from flask import Flask, request

app = Flask(__name__)

@app.route('/user', methods=['GET'])
def greet_user():
    name = request.args.get('name')
    email = request.args.get('email')

    if name and email:
        greeting = f"Hello, {name}! Your email is {email}."
    else:
        greeting = "Please provide both name and email parameters."

    return greeting

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
