from flask_wtf import FlaskForm
from wtforms import StringField, SubmitField, TextAreaField
from wtforms.fields.html5 import DateField
from wtforms.validators import Length, Email, DataRequired


class FutureEmail(FlaskForm):
    subject = StringField(label='Subject', validators=[Length(0, 128)])
    email = StringField(label='Email Address', validators=[DataRequired(), Length(0, 64), Email()])
    date = DateField('Date to receive this email', validators=[DataRequired()], format='%Y/%m/%d')
    content = TextAreaField('Content')
    submit = SubmitField('Submit')
