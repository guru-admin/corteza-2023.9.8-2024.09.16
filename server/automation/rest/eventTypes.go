package rest

import (
	"context"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/rest/request"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/automation/service"
)

type (
	EventTypes struct {
		reg interface {
			Types() []string
		}
	}

	eventTypePayload struct {
		Set []eventTypeDef `json:"set"`
	}

	eventTypeDef struct {
		ResourceType string                   `json:"resourceType"`
		EventType    string                   `json:"eventType"`
		Properties   []eventTypePropertyDef   `json:"properties"`
		Constraints  []eventTypeConstraintDef `json:"constraints"`
	}

	eventTypePropertyDef struct {
		Name      string `json:"name"`
		Type      string `json:"type"`
		Immutable bool   `json:"immutable"`
	}

	eventTypeConstraintDef struct {
		Name string `json:"name"`
	}
)

func (EventTypes) New() *EventTypes {
	ctrl := &EventTypes{reg: service.Registry()}
	return ctrl
}

func (ctrl EventTypes) List(_ context.Context, _ *request.EventTypesList) (interface{}, error) {
	return eventTypePayload{Set: getEventTypeDefinitions()}, nil
}
