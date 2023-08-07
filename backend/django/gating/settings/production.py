from gating.settings.common import *
from dotenv import load_dotenv
import os
import dj_database_url

BASE_DIR = Path(__file__).resolve().parent.parent.parent

load_dotenv('%s/.env' % BASE_DIR)

# SECURITY WARNING: keep the secret key used in production secret!
SECRET_KEY = os.environ.get('DJANGO_KEY')

# SECURITY WARNING: don't run with debug turned on in production!
DEBUG = True

# Database
# https://docs.djangoproject.com/en/3.2/ref/settings/#databases
if "DATABASE_URL" in os.environ:
    DATABASES = {}
    DATABASES['default'] = dj_database_url.config(conn_max_age=600, env="DATABASE_URL")
else:
    DATABASES = {
        'default': {
            'ENGINE': 'django.db.backends.postgresql',
            'NAME': 'gating',
            'USER': 'postgres',
            'HOST': '127.0.0.1',
            'PASSWORD': os.environ.get('USER_PASSWORD'),
        }
    }

# ENABLE_2FA = True
# ALLOWED_HOSTS = ['.gating.com']
