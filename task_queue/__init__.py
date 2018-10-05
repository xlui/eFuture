# Redis version delayed task queue implement
import redis

from config import redisHost, redisPort, redisPassword, redisDB
from .redis_impl import push, pop

QUEUE_KEY = 'mail_queue'
pool = redis.ConnectionPool(host=redisHost, port=redisPort, db=redisDB, password=redisPassword)
connection = redis.Redis(connection_pool=pool)
# delete the following line when deploy
connection.delete(QUEUE_KEY)
