from flask import render_template, request, redirect, url_for

from . import main
from .forms import FutureEmail


@main.route('/', methods=['GET', 'POST'])
def index():
    form = FutureEmail()
    if request.method == 'POST':
        print(form.subject)
        print(form.email)
        print(form.date)
        print(form.content)
        return redirect(url_for('main.index'))
    return render_template('index.html', form=form)
