# Generated by Django 3.2.3 on 2023-05-22 02:57

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('gatingDB', '0002_auto_20230521_2122'),
    ]

    operations = [
        migrations.AlterField(
            model_name='awaitingconnection',
            name='interactionAppID',
            field=models.BigIntegerField(),
        ),
        migrations.AlterField(
            model_name='awaitingconnection',
            name='interactionToken',
            field=models.BigIntegerField(),
        ),
    ]