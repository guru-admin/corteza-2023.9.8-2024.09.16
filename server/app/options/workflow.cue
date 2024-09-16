package options

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)

workflow: schema.#optionsGroup & {
	handle: "workflow"
	options: {
		register: {
			type:          "bool"
			defaultGoExpr: "true"
			description:   "Registers enabled and valid workflows and executes them when triggered"
		}
		exec_debug: {
			type:        "bool"
			description: "Enables verbose logging for workflow execution"
		}
		call_stack_size: {
			type:          "int"
			defaultGoExpr: "16"
			description:   "Defines the maximum call stack size between workflows"
		}
		stack_trace_enabled: {
			type:          "bool"
			defaultGoExpr: "true"
			description:   "Enables execution stack trace construction"
		}
		stack_trace_full: {
			type:          "bool"
			defaultGoExpr: "true"
			description:   "Forces the stack trace to record all steps"
		}
	}
	title: "Workflow"
}
