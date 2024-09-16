package rest

import (
	"github.com/go-chi/chi/v5"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/rest/handlers"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/auth"
)

func MountRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		// Protect all _private_ routes
		r.Group(func(r chi.Router) {
			r.Use(auth.HttpTokenValidator("api"))

			handlers.NewWorkflow(Workflow{}.New()).MountRoutes(r)
			handlers.NewTrigger(Trigger{}.New()).MountRoutes(r)
			handlers.NewSession(Session{}.New()).MountRoutes(r)
			handlers.NewFunction(Function{}.New()).MountRoutes(r)
			handlers.NewType(Type{}.New()).MountRoutes(r)
			handlers.NewPermissions(Permissions{}.New()).MountRoutes(r)
			handlers.NewEventTypes(EventTypes{}.New()).MountRoutes(r)
		})
	}
}
