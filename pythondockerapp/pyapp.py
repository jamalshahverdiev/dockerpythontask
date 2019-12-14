from flask import Flask

app = Flask(__name__)

@app.route('/', methods = ['GET'])
def mainroute():
    return "Success!"

@app.route('/ping', methods = ['GET'])
def pintGroute():
    return "Ok"

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080, debug=True)
