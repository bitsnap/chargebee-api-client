package enums

type SourceEnum string

const (
	AdminConsoleSource    SourceEnum = "admin_console"
	APISource                        = "api"
	ScheduledJobSource               = "scheduled_job"
	HostedPageSource                 = "hosted_page"
	PortalSource                     = "portal"
	SystemSource                     = "system"
	NoneSource                       = "none"
	JSAPISource                      = "js_api"
	MigrationSource                  = "migration"
	BulkOperationSource              = "bulk_operation"
	ExternalServiceSource            = "external_service"
)
