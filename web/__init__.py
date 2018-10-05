from flask import Flask
from flask_bootstrap import Bootstrap


def create_app():
    bootstrap = Bootstrap()
    app = Flask(__name__)
    app.config.update({
        'SECRET_KEY': 'hard to guess string'
    })

    bootstrap.init_app(app)

    from .main import main
    app.register_blueprint(main)

    return app
