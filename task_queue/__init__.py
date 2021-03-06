# Redis version delayed task queue implement
import redis

from config import redisHost, redisPort, redisPassword, redisDB

QUEUE_KEY = 'mail_queue'
pool = redis.ConnectionPool(host=redisHost, port=redisPort, db=redisDB, password=redisPassword)
connection = redis.Redis(connection_pool=pool)
# delete the following line when deploy
# connection.delete(QUEUE_KEY)

# To avoid circulate dependency
from .redis_impl import push, pop
