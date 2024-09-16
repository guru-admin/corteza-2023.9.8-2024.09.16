package app

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/app/options"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/compose"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/federation"
)

corteza: schema.#platform & {
	"ident": "corteza"

	"options": [
		options.DB,
		options.HTTPClient,
		options.HTTPServer,
		options.RBAC,
		options.SCIM,
		options.SMTP,
		options.actionLog,
		options.apigw,
		options.auth,
		options.corredor,
		options.environment,
		options.eventbus,
		options.federation,
		options.limit,
		options.locale,
		options.log,
		options.messagebus,
		options.monitor,
		options.objectStore,
		options.provision,
		options.sentry,
		options.template,
		options.upgrade,
		options.waitFor,
		options.websocket,
		options.workflow,
		options.discovery,
		options.attachment,
		options.webapp,
	]

	// platform resources
	"resources": resources

	"components": [
		system.component,
		compose.component,
		automation.component,
		federation.component,
	]
}
