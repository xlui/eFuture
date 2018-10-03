import smtplib
from email.mime.text import MIMEText
from email.header import Header
from config.config import configuration

charset = 'utf-8'
host = configuration['smtp']
port = configuration['smtp_port']
username = configuration['username']
password = configuration['password']
sender = username
receivers = ['liuqi0315@gmail.com']

message = MIMEText('This email is sent from python', _charset=charset)
message['From'] = Header(sender, charset)
message['To'] = Header(', '.join(receivers), charset)
message['Subject'] = Header('eFuture 测试邮件', charset)


print(host)
try:
    smtp = smtplib.SMTP(host=host, port=port)
    smtp.ehlo()
    smtp.starttls()
    smtp.login(username, password)
    smtp.sendmail(sender, receivers, message.as_string())
    print('Successfully send')
except smtplib.SMTPException as e:
    print('Cannot send email:', e)
