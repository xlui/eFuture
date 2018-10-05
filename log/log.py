import logging
from logging.handlers import RotatingFileHandler

from config import logPath, logLevel

logger = logging.getLogger('eFuture')
formatter = logging.Formatter('%(asctime)s - %(filename)s:%(lineno)s - [%(levelname)-5.5s] - %(message)s')

# 10MB each log file
fileHandler = RotatingFileHandler(logPath, maxBytes=10 * 1024 * 1024, backupCount=5)
fileHandler.setFormatter(formatter)
logger.addHandler(fileHandler)

# output to console
consoleHandler = logging.StreamHandler()
consoleHandler.setFormatter(formatter)
logger.addHandler(consoleHandler)

logger.setLevel(logLevel)
