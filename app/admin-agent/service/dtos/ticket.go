package dtos

import (
	"fmt"
	"go-admin/app/admin-agent/model"
	"go-admin/common/dto"
)

type TicketPagesReq struct {
	dto.Pagination
	Type   string `json:"type"`
	Status string `json:"status"`
}

func (r *TicketPagesReq) GenConditions() string {
	switch {
	case len(r.Type) != 0 && len(r.Status) != 0:
		return fmt.Sprintf("type = %s AND status = %s", r.Type, r.Status)
	case len(r.Type) != 0:
		return fmt.Sprintf("type = %s", r.Type)
	case len(r.Status) != 0:
		return fmt.Sprintf("status = %s", r.Status)
	default:
		return ""
	}
}

type TicketReq struct {
	RelaID     int    `json:"relaID"`
	Name       string `json:"name"`
	Properties string `json:"properties"`
	Type       string `json:"type"`
	UserID     int    `json:"userID"`
	OptMsg     string `json:"optMsg"`
}

func (r *TicketReq) Generate(model *model.Ticket) {
	if r.RelaID != 0 {
		model.RelaID = r.RelaID
	}
	if len(r.Name) != 0 {
		model.Name = r.Name
	}
	if len(r.Properties) != 0 {
		model.Properties = r.Properties
	}
	if len(r.Type) != 0 {
		model.Type = r.Type
	}
}
