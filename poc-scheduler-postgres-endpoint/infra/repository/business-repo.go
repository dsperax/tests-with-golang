package repository

import domainNot "poc-scheduler-postgres-endpoint/domain/notification"

type IBusinessRepository interface {
	EndPointNotificacaoEmails(notificationEntity domainNot.NotificationEntity) error
}
