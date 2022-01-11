package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/pkg/errors"
)

// Response and Query Definitions

// Client Definition

type AuditLogClient interface {
  Get(ctx context.Context, params *AuditLogQueryParams) (*AuditLogsResponse, error)
}


type RESTAuditLogClient struct {
  client *APIClient
}

var _ AuditLogClient = &RESTAuditLogClient{}

func (c *RESTAuditLogClient) restClient() *sling.Sling {
  return c.client.client()
}

func (c *RESTAuditLogClient) Get(ctx context.Context, params *AuditLogQueryParams) (*AuditLogsResponse, error) {
  res := &AuditLogsResponse{}
  resp, err := c.restClient().Get("auditlog").Receive(res, nil)
  if err != nil {
    return nil, errors.Wrap(err, "could not get audit logs")
  }

  if resp.StatusCode != http.StatusOK {
    return nil, fmt.Errorf("could not get audit logs: %s", resp.Status)
  }
  
  return res, nil
}

