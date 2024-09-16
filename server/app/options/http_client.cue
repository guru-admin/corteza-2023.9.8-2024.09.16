package options

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

HTTPClient: schema.#optionsGroup & {
	title: "HTTP Client"
	// Explicitly define all variants to be 100% compaltible with old name
	handle: "http-client"

	// @todo remove explcitly defined expIdent and adjust the code
	expIdent: "HTTPClient"

	imports: [
		"\"time\"",
	]

	options: {
		tls_insecure: {
			type: "bool"
			description: """
				Allow insecure (invalid, expired TLS/SSL certificates) connections.

				[IMPORTANT]
				====
				We strongly recommend keeping this value set to false except for local development or demos.
				====
				"""
		}
		timeout: {
			type:        "time.Duration"
			description: "Default timeout for clients."

			defaultGoExpr: "30 * time.Second"
			defaultValue:  "30s"
		}
	}
}
