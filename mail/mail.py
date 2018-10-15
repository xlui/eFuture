import smtplib
from email.header import Header
from email.mime.text import MIMEText

from config import host, port, username, password
from log import logger

__charset = 'utf-8'


def send_mail(receiver: str, subject: str, message: str):
    message = MIMEText(message, _charset=__charset)
    message['From'] = Header('Future Email <{}>'.format(username))
    message['To'] = Header('<{}>'.format(receiver))
    message['Subject'] = Header(subject)

    try:
        smtp = smtplib.SMTP(host=host, port=port)
        smtp.starttls()
        smtp.login(username, password)
        smtp.sendmail(username, [receiver], message.as_string())
        logger.info('Successfully send')
    except smtplib.SMTPException as e:
        logger.error('Cannot send email: {0}'.format(e))


if __name__ == '__main__':
    send_mail('liuqi0315@gmail.com', 'Test send mail to gmail', 'Sent from python')
