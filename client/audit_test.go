package client

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAuditLog(t *testing.T) {

  t.Log("Testing audit log get client")

  // Get time in unix from 5 minutes ago
  startTime := time.Now().Add(time.Duration(-5) * time.Minute).Unix()
  endTime := time.Now().Unix()

  testQueryParams := &AuditLogQueryParams{
    StartTime: int32(startTime),
    EndTime: int32(endTime),
    PageNumber: 1,
    PerPage: 100,
  }

  ac, _ := NewRestClientFromConfigFile()
  res, err := ac.AuditLog().Get(context.TODO(), testQueryParams)
  t.Log(res)

  out, _ := os.Create("audit_test.json")
  defer out.Close()

  json.NewEncoder(out).Encode(res)

  require.NoError(t, err)

  assert.NotNil(t, res)
}