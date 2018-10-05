import json
import os

DEFAULT_PATH = '/data/eFuture/config.json'
file_path = os.getenv('EFUTURE_CONFIG', DEFAULT_PATH)
with open(file_path, encoding='utf-8', mode='r') as f:
    configuration = json.load(f)
# print(configuration)

host = configuration['smtp']
port = configuration['smtp_port']
username = configuration['username']
password = configuration['password']
redisHost = configuration['redis_host']
redisPort = configuration['redis_port']
redisPassword = configuration['redis_password']
redisDB = configuration['redis_db']
logPath = configuration['log_path']
logLevel = configuration['log_level']
