package options

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

websocket: schema.#optionsGroup & {
	handle: "websocket"

	imports: [
		"\"time\"",
	]
	title: "Websocket server"

	options: {
		log_enabled: {
			type:        "bool"
			description: "Enable extra logging for authentication flows"
		}
		timeout: {
			type:        "time.Duration"
			description: "Time before `WsServer` gets timed out."

			defaultGoExpr: "15 * time.Second"
			defaultValue:  "15s"
		}
		ping_timeout: {
			type:          "time.Duration"
			defaultGoExpr: "120 * time.Second"
			defaultValue:  "120s"
		}
		ping_period: {
			type: "time.Duration"

			defaultGoExpr: "((120 * time.Second) * 9) / 10"
			defaultValue:  "119s"
		}
	}
}
