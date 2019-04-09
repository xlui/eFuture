import datetime

from flask import render_template, redirect, url_for

from config import Letter
from log import logger
from task_queue import push
from . import main
from .forms import FutureEmail


@main.route('/', methods=['GET', 'POST'])
def index():
    form = FutureEmail()
    if form.validate_on_submit():
        letter = Letter(
            form.subject.data,
            form.receiver.data,
            form.receiveDate.data.__str__(),  # type: datetime.date
            form.content.data,
            datetime.datetime.today().strftime('%Y-%m-%d %H:%M:%S')
        )
        push(letter.toJSON(), datetime.datetime.strptime(letter.receiveDate, '%Y-%m-%d'))
        logger.info(f'Receive new future email! Letter: {letter}')
        return redirect(url_for('main.index'))
    logger.info(f'New request or invalid request!')
    return render_template('index.html', form=form)
