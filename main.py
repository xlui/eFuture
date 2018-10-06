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
        from task_queue import pop
        import json
        while True:
            time.sleep(1)
            boolean, message = pop()
            if boolean:
                letter = json.loads(message)  # type:dict
                print(letter['subject'])
                print(letter['sender'])
                print(letter['receiveDate'])
                print(letter['sendDate'])
                print(letter['content'])


    thread = threading.Thread(target=run, args=())
    thread.setDaemon(True)
    thread.start()
    manager.add_command('shell', Shell(make_context=make_shell_context))
    manager.run()
