import datetime
import uuid

from task_queue import connection, QUEUE_KEY


def push(message: str, date: datetime.datetime):
    """Push a message into redis zset

    :param message: message content
    :param date: the date this message to be consumed
    :return: None
    """
    msg_id = uuid.uuid4()
    pipeline = connection.pipeline()
    pipeline.set(msg_id, message)
    pipeline.zadd(QUEUE_KEY, msg_id, date.timestamp())
    pipeline.execute()


def pop():
    """Check the first task in redis(which is the task with the smallest score)
    if the score(timestamp) is smaller or equal to current timestamp, the task
    should be take out and done.

    :return: True if task is take out, and False if it is not the time.
    """
    task = connection.zrange(QUEUE_KEY, 0, 0)
    if not task:
        return False, ""
    msg_id = task[0]
    timestamp = connection.zscore(QUEUE_KEY, msg_id)
    now = datetime.datetime.now().timestamp()
    if timestamp < now or abs(timestamp - now) <= 1e-6:
        message = connection.get(msg_id)
        pipeline = connection.pipeline()
        pipeline.zrem(QUEUE_KEY, msg_id)
        pipeline.delete(msg_id)
        pipeline.execute()
        return True, message
    return False, ""


if __name__ == '__main__':
    now = datetime.datetime.now()
    from log import logger

    logger.debug('push hello')
    push('hello', now + datetime.timedelta(seconds=10))
    while True:
        b, m = pop()
        if b:
            logger.debug(m)
