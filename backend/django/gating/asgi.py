"""
ASGI config for gating project.

It exposes the ASGI callable as a module-level variable named ``application``.

For more information on this file, see
https://docs.djangoproject.com/en/3.2/howto/deployment/asgi/
"""

import os

from django.core.asgi import get_asgi_application
from pathlib import Path
from dotenv import load_dotenv

application = get_asgi_application()

load_dotenv('%s/.env' % Path(__file__).resolve().parent.parent.parent)

os.environ.setdefault('DJANGO_SETTINGS_MODULE', f'gating.settings.{os.environ.get("ENV")}')

print(load_dotenv('%s/.env' % Path(__file__).resolve().parent.parent.parent),f'gating.settings.{os.environ.get("ENV")}')