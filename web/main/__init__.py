from flask import Blueprint

main = Blueprint('main', __name__)

# To avoid cyclic dependence
from . import views
