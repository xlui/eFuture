import logging
from logging.handlers import RotatingFileHandler
from config.config import logPath, logLevel

# 10MB each log file
handler = RotatingFileHandler(logPath, maxBytes=10 * 1024 * 1024, backupCount=5)
fmt = '%(asctime)s - %(filename)s:%(lineno)s - func: [%(name)s] - %(message)s'
formatter = logging.Formatter(fmt)
handler.setFormatter(formatter)

logger = logging.getLogger('eFuture')
logger.addHandler(handler)
logger.setLevel(logLevel)
logger.debug('debug')
