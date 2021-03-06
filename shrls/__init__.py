#!/usr/bin/env python
# -*- coding: utf-8 -*-

import os
import random

from flask import Flask
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__, static_url_path='/static')
app.config['SECRET_KEY'] = str(random.random())
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:////%s/urls.db' % os.getcwd()
app.config['UPLOAD_FOLDER'] = '%s/uploads/' % os.getcwd()

app.config['shrls_username'] = 'admin'
app.config['shrls_password'] = 'changemenow'

app.config['shrls_redirect_unknown'] = True
app.config['shrls_redirect_url'] = 'http://example.com/'
app.config['shrls_base_url'] = 'http://example.com/'

if not os.path.exists(app.config['UPLOAD_FOLDER']):
    os.makedirs(app.config['UPLOAD_FOLDER'])

import shrls.views
from shrls.models import DBSession

@app.teardown_appcontext
def shutdown_session(exception=None):
    DBSession.remove()
