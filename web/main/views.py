import datetime
from flask import render_template, request, redirect, url_for

from config import Letter
from . import main
from .forms import FutureEmail


@main.route('/', methods=['GET', 'POST'])
def index():
    form = FutureEmail()
    if request.method == 'POST':
        letter = Letter(
            form.subject.data,
            form.email.data,
            form.date.data.__str__(),
            form.content.data,
            datetime.datetime.today().strftime('2018-10-06 11:03:19')
        )
        print(letter)
        return redirect(url_for('main.index'))
    return render_template('index.html', form=form)
