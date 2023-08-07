from django.contrib import admin
from . import models


class CustomUserAdmin(admin.ModelAdmin):
    raw_id_fields = ['user']
    search_fields = ['user']


class WalletInline(admin.TabularInline):
    model = models.wallet
    extra = 0


class DiscordUserAdmin(admin.ModelAdmin):
    list_display = ['id', 'name', 'createdAt']
    search_fields = ['id', 'name']
    inlines = [WalletInline]


class GuildAdmin(admin.ModelAdmin):
    list_display = ['id','createdAt']
    search_fields = ['id']


class NetworkAdmin(admin.ModelAdmin):
    list_display = ['chainID', 'name', 'httpRPC', 'createdAt']
    search_fields = ['chainID','name']


class WalletAdmin(admin.ModelAdmin):
    list_display = ['id', 'address', 'network', 'user', 'createdAt']
    search_fields = ['id', 'address']


class ContractAdmin(admin.ModelAdmin):
    list_display = ['id', 'address', 'type', 'network', 'createdAt']
    search_fields = ['id', 'address', 'type']


class RoleAdmin(admin.ModelAdmin):
    list_display = ['id', 'guild', 'createdAt']
    search_fields = ['id', 'guild']
    raw_id_fields = ['contracts']

class ContractRoleAdmin(admin.ModelAdmin):
    list_display = ['role', 'contract']
    search_fields = ['role']

class AwaitingConnectionAdmin(admin.ModelAdmin):
    list_display = ['message', 'user', 'guild', 'interactionToken', 'interactionAppID', 'interactionID', 'createdAt']
    search_fields = ['message', 'user', 'guild']

admin.site.register(models.django_user, CustomUserAdmin)
admin.site.register(models.user, DiscordUserAdmin)
admin.site.register(models.guild, GuildAdmin)
admin.site.register(models.network, NetworkAdmin)
admin.site.register(models.wallet, WalletAdmin)
admin.site.register(models.contract, ContractAdmin)
admin.site.register(models.role, RoleAdmin)
admin.site.register(models.contract_role, ContractRoleAdmin)
admin.site.register(models.awaitingConnection, )
