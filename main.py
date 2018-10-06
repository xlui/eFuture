import threading

from web import create_app

app = create_app()

if __name__ == '__main__':
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
    app.run(host='0.0.0.0', port=5000)
