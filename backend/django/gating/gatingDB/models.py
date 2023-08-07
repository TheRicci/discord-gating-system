from django.db import models
from django.contrib.auth.models import AbstractUser

class django_user(AbstractUser):
    user = models.ForeignKey('user', blank=True, null=True, on_delete=models.CASCADE)

    class Meta:
        verbose_name_plural = "Django Users"
class user(models.Model):
    id = models.BigIntegerField(primary_key=True)
    name = models.CharField(max_length=255, blank=True, null=True)
    createdAt = models.DateTimeField()

class guild(models.Model):
    id = models.BigIntegerField(primary_key=True)
    name = models.CharField(max_length=255, blank=True, null=True)
    createdAt = models.DateTimeField()

class network(models.Model):
    chainID = models.IntegerField(primary_key=True)
    name = models.CharField(max_length=255, blank=True, null=True)
    httpRPC = models.URLField(blank=True, null=True)
    createdAt = models.DateTimeField()

class wallet(models.Model):
    address = models.CharField(max_length=255)
    network = models.ForeignKey('network', on_delete=models.CASCADE)
    user = models.ForeignKey('user', on_delete=models.CASCADE)
    createdAt = models.DateTimeField()

    class Meta:
        constraints = [
            models.UniqueConstraint(fields=['address', 'network'], name='wallet_unique_address_network')
        ]

class contract(models.Model):
    address = models.CharField(max_length=255)
    type = models.IntegerField()
    network = models.ForeignKey('network', on_delete=models.CASCADE)
    createdAt = models.DateTimeField()

    class Meta:
        constraints = [
            models.UniqueConstraint(fields=['address', 'network'], name='contract_unique_address_network')
        ]


class role(models.Model):
    id = models.BigIntegerField(primary_key=True)
    guild = models.ForeignKey('guild', on_delete=models.CASCADE)
    createdAt = models.DateTimeField()
    contracts = models.ManyToManyField(
        contract,
        related_name='roles',
        through='contract_role'
    )

class contract_role(models.Model):
    role = models.ForeignKey('role', on_delete=models.CASCADE, related_name='role_contracts')
    contract = models.ForeignKey('contract', on_delete=models.CASCADE, related_name='role_contracts')

    class Meta:
        constraints = [
            models.UniqueConstraint(fields=['role', 'contract'], name='contract_role_unique_role_contract')
        ]
        verbose_name_plural = "Contracts and Roles"

class awaitingConnection(models.Model):
    message = models.TextField(primary_key=True)
    user = models.ForeignKey('user', on_delete=models.CASCADE)
    guild = models.ForeignKey('guild', on_delete=models.CASCADE)
    interactionID = models.BigIntegerField()
    interactionToken = models.BigIntegerField()
    interactionAppID = models.BigIntegerField()
    createdAt = models.DateTimeField()

