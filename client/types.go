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

/**
 Notifications
*/

// TODO: Populate when I can view sample data
type NotificationsResponse map[string]interface{}

/**
	Cameras
*/

// Get All Camers
type CamerasResponse struct {
	Cameras []CameraResponse `json:"cameras"`
}

type CameraResponse struct {
	CameraId string `json:"camera_id"`
	CloudRetention int `json:"cloud_retention"`
	DateAdded int `json:"data_added"`
	DeviceRetention int `json:"device_retention"`
	Firmware string `json:"firmware"`
	LastOnline int `json:"last_online"`
	LocalIp string `json:"local_ip"`
	Location string `json:"location"`
	LocationAngle float32 `json:"location_angle"`
	LocationLat float32 `json:"location_lat"`
	LocationLon float32 `json:"location_lon"`
	MacAddress string `json:"mac"`
	Model	string	`json:"model"`
	Name	string `json:"name"`
	Serial string `json:"serial"`
	Site string `json:"site"`
	Status string `json:"status"`
}