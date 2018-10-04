import redis

pool = redis.ConnectionPool(host='localhost', port=6379)
r = redis.Redis(connection_pool=pool)
r.set('a', 'the value of key a')
print(r.get('a'))
print(r['a'])
print(type(r.get('a')))
