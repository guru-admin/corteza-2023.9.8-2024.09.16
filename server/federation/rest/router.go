package rest

import (
	"github.com/go-chi/chi/v5"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/federation/rest/handlers"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/auth"
)

func MountRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Group(func(r chi.Router) {
			handlers.NewNodeHandshake(NodeHandshake{}.New()).MountRoutes(r)
		})

		// Protect all _private_ routes
		r.Group(func(r chi.Router) {
			r.Use(auth.HttpTokenValidator("api"))

			handlers.NewPermissions(Permissions{}.New()).MountRoutes(r)

			handlers.NewNode(Node{}.New()).MountRoutes(r)
			handlers.NewManageStructure((ManageStructure{}.New())).MountRoutes(r)

			handlers.NewSyncData((SyncData{}.New())).MountRoutes(r)
			handlers.NewSyncStructure((SyncStructure{}.New())).MountRoutes(r)
		})
	}
}
