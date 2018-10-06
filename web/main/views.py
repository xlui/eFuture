import datetime

from flask import render_template, request, redirect, url_for

from config import Letter
from task_queue import push
from . import main
from .forms import FutureEmail


@main.route('/', methods=['GET', 'POST'])
def index():
    form = FutureEmail()
    if request.method == 'POST':
        letter = Letter(
            form.subject.data,
            form.email.data,
            form.date.data.__str__(),  # type: datetime.date
            form.content.data,
            datetime.datetime.today().strftime('%Y-%m-%d %H:%M:%S')
        )
        push(str(letter), datetime.datetime.strptime(letter.receiveDate, '%Y-%m-%d'))
        return redirect(url_for('main.index'))
    return render_template('index.html', form=form)
