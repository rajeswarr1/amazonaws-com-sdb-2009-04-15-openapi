package main

import (
	"github.com/amazon-simpledb/mcp-server/config"
	"github.com/amazon-simpledb/mcp-server/models"
	tools_action_batchdeleteattributes "github.com/amazon-simpledb/mcp-server/tools/action_batchdeleteattributes"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_action_batchdeleteattributes.CreateGet_batchdeleteattributesTool(cfg),
	}
}
