import datetime

from log import logger
from task_queue import push, pop

if __name__ == '__main__':
    logger.debug('push hello')
    push('hello', datetime.datetime.now() + datetime.timedelta(seconds=10))
    while True:
        try:
            b, msg = pop()
            if b:
                logger.debug('pop hello')
        except KeyboardInterrupt as e1:
            break
        except Exception as e:
            print(e)
