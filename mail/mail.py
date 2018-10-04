import smtplib
from email.header import Header
from email.mime.text import MIMEText

from config.config import host, port, username, password

__charset = 'utf-8'


def send_mail(receivers, subject, message):
    message = MIMEText(message, _charset=__charset)
    message['From'] = Header('Future Email', __charset)
    message['To'] = Header(', '.join(receivers), __charset)
    message['Subject'] = Header(subject, __charset)

    try:
        smtp = smtplib.SMTP(host=host, port=port)
        smtp.starttls()
        smtp.login(username, password)
        smtp.sendmail(username, receivers, message.as_string())
        print('Successfully send')
    except smtplib.SMTPException as e:
        print('Cannot send email:', e)


if __name__ == '__main__':
    from datetime import datetime
    send_mail(['liuqi0315@gmail.com'], 'This is also a test email', 'Sent from python!\nAnd current is ' + datetime.now().strftime('2018-10-4 14:20:20'))
