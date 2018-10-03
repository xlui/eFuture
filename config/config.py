import os
import json

file_path = os.getenv('EFUTURE_CONFIG', '/data/eFuture/config.json')
with open(file_path, encoding='utf-8', mode='r') as f:
    configuration = json.load(f)
# print(configuration)
