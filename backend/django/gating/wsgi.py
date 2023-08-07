"""
WSGI config for gating project.

It exposes the WSGI callable as a module-level variable named ``application``.

For more information on this file, see
https://docs.djangoproject.com/en/3.2/howto/deployment/wsgi/
"""

import os

from django.core.wsgi import get_wsgi_application
from pathlib import Path
from dotenv import load_dotenv

application = get_wsgi_application()

load_dotenv('%s/.env' % Path(__file__).resolve().parent.parent.parent)

os.environ.setdefault('DJANGO_SETTINGS_MODULE', f'gating.settings.{os.environ.get("ENV")}')

print(load_dotenv('%s/.env' % Path(__file__).resolve().parent.parent.parent),f'gating.settings.{os.environ.get("ENV")}')