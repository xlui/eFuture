from flask_script import Manager, Shell

from web import create_app

app = create_app()
manager = Manager(app)

if __name__ == '__main__':
    def make_shell_context():
        return dict(app=app, )


    manager.add_command('shell', Shell(make_context=make_shell_context))
    manager.run()
