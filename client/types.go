/*
  Type Definitions for QueryParams, Response Objects, etc.
**/

package client

/**
AuditLog Types
*/

type AuditLogQueryParams struct {
	StartTime  int32 `url:"start_time,omitempty"`
	EndTime    int32 `url:"end_time,omitempty"`
	PageNumber int   `url:"page_num,omitempty"`
	PerPage    int   `url:"per_page,omitempty"`
}

type AuditLogsResponse map[string]interface{}

// TODO: Need to figure out why this is not being parsed by Sling
// type AuditLogsResponse struct {
// 	AuditLogs []AuditLogResponse `json:"audit_logs"`
// 	HasNext   bool               `json:"has_next"`
// }

type AuditLogResponse struct {
	Details          string           `json:"details"`
	Devices          []DeviceResponse `json:"devices"`
	EventDescription string           `json:"event_description"`
	EventName        string           `json:"event_name"`
	IpAddress        string           `json:"ip_address"`
	OrganizationId   string           `json:"organization_id"`
	TimeStamp        string           `json:"timestamp"`
	UserEmail        string           `json:"user_email"`
	UserId           string           `json:"user_id"`
	UserName         string           `json:"user_name"`
	VerkadaSupportId string           `json:"verkada_support_id"`
}

type DeviceResponse struct {
	Details        string `json:"details"`
	DeviceId       string `json:"device_id"`
	DeviceName     string `json:"device_name"`
	DeviceSiteName string `json:"device_site_name"`
	DeviceType     string `json:"device_time"`
}
