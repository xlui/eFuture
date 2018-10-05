from flask import render_template, redirect, url_for

from . import main
from .forms import FutureEmail


@main.route('/', methods=['GET', 'POST'])
def index():
    form = FutureEmail()
    if form.validate_on_submit():
        print(form.subject)
        print(form.email)
        print(form.date)
        print(form.content)
        return redirect(url_for('main.index'))
    return render_template('index.html', form=form)
