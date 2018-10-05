from flask_wtf import FlaskForm
from wtforms import StringField, DateField, SubmitField
from wtforms.validators import Length, Email, DataRequired


class FutureEmail(FlaskForm):
    subject = StringField(label='Subject', validators=[Length(0, 128)])
    email = StringField(label='Email Address', validators=[DataRequired(), Length(0, 64), Email()])
    date = DateField(label='Date to receive this email')
    content = StringField('Content')
    submit = SubmitField('Submit')
