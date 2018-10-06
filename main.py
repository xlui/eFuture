import threading

from flask_script import Manager, Shell

from web import create_app

app = create_app()
manager = Manager(app)

if __name__ == '__main__':
    def make_shell_context():
        return dict(app=app, )


    def run():
        import time
        import json
        from task_queue import pop
        from mail import send_mail
        while True:
            time.sleep(1)
            boolean, message = pop()
            if boolean:
                letter = json.loads(message)  # type:dict
                send_mail(letter['receiver'], letter['subject'], letter['content'])


    thread = threading.Thread(target=run, args=())
    thread.setDaemon(True)
    thread.start()
    manager.add_command('shell', Shell(make_context=make_shell_context))
    manager.run()
