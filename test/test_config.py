import re
import unittest
from config.config import configuration


class TestConfig(unittest.TestCase):
    def __init__(self, methodName='runTest'):
        self.email_regexp = re.compile(r'^\w+[-.]?\w+@\w+[-]?\w+\.\w+$')
        super().__init__(methodName)

    def testUsername(self):
        self.assertTrue(self.email_regexp.match(configuration['username']), msg='Email username invalid!!')

    def testPassword(self):
        self.assertNotEqual("", configuration['password'], msg='Admin email password must not be empty!')

    def testSMTP(self):
        self.assertNotEqual("", configuration['smtp'], msg='SMTP server should not be empty!')

    def testSMTPPort(self):
        self.assertRegex(configuration['smtp_port'], '^\d+$', msg='SMTP port must be numbers!')
