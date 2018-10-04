import json
import os

file_path = os.getenv('EFUTURE_CONFIG', '/data/eFuture/config.json')
with open(file_path, encoding='utf-8', mode='r') as f:
    configuration = json.load(f)
# print(configuration)

host = configuration['smtp']
port = configuration['smtp_port']
username = configuration['username']
password = configuration['password']
logPath = configuration['log_path']
logLevel = configuration['log_level']