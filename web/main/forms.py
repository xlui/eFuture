from datetime import datetime

from flask_wtf import FlaskForm
from wtforms import StringField, SubmitField, TextAreaField
from wtforms.fields.html5 import DateField
from wtforms.validators import Length, Email, DataRequired


class FutureEmail(FlaskForm):
    subject = StringField(label='Subject', validators=[Length(0, 128)], default='A Future Email')
    receiver = StringField(label='Email Address', validators=[DataRequired(), Length(0, 64), Email()],
                           default='liuqi0315@gmail.com')
    receiveDate = DateField('Date to receive this email', validators=[DataRequired()], format='%Y-%m-%d',
                            default=datetime.today)
    content = TextAreaField('Content', default="Hi there, I'm now writing to you at {}.".format(
        datetime.now().strftime('%Y-%m-%d %H:%M:%S')))
    submit = SubmitField('Submit')
