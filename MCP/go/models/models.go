package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Attribute represents the Attribute schema from the OpenAPI specification
type Attribute struct {
	Value interface{} `json:"Value"`
	Alternatenameencoding interface{} `json:"AlternateNameEncoding,omitempty"`
	Alternatevalueencoding interface{} `json:"AlternateValueEncoding,omitempty"`
	Name interface{} `json:"Name"`
}

// InvalidQueryExpression represents the InvalidQueryExpression schema from the OpenAPI specification
type InvalidQueryExpression struct {
	Boxusage float32 `json:"BoxUsage,omitempty"`
}

// ReplaceableAttribute represents the ReplaceableAttribute schema from the OpenAPI specification
type ReplaceableAttribute struct {
	Name interface{} `json:"Name"`
	Replace interface{} `json:"Replace,omitempty"`
	Value interface{} `json:"Value"`
}

// UpdateCondition represents the UpdateCondition schema from the OpenAPI specification
type UpdateCondition struct {
	Exists interface{} `json:"Exists,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// ReplaceableItem represents the ReplaceableItem schema from the OpenAPI specification
type ReplaceableItem struct {
	Attributes interface{} `json:"Attributes"`
	Name interface{} `json:"Name"`
}

// TooManyRequestedAttributes represents the TooManyRequestedAttributes schema from the OpenAPI specification
type TooManyRequestedAttributes struct {
	Boxusage float32 `json:"BoxUsage,omitempty"`
}

// NumberDomainsExceeded represents the NumberDomainsExceeded schema from the OpenAPI specification
type NumberDomainsExceeded struct {
	Boxusage float32 `json:"BoxUsage,omitempty"`
}

// GetAttributesResult represents the GetAttributesResult schema from the OpenAPI specification
type GetAttributesResult struct {
	Attributes interface{} `json:"Attributes,omitempty"`
}

// BatchDeleteAttributesRequest represents the BatchDeleteAttributesRequest schema from the OpenAPI specification
type BatchDeleteAttributesRequest struct {
	Items interface{} `json:"Items"`
	Domainname interface{} `json:"DomainName"`
}

// RequestTimeout represents the RequestTimeout schema from the OpenAPI specification
type RequestTimeout struct {
	Boxusage float32 `json:"BoxUsage,omitempty"`
}

// GetAttributesRequest represents the GetAttributesRequest schema from the OpenAPI specification
type GetAttributesRequest struct {
	Attributenames interface{} `json:"AttributeNames,omitempty"`
	Consistentread interface{} `json:"ConsistentRead,omitempty"`
	Domainname interface{} `json:"DomainName"`
	Itemname interface{} `json:"ItemName"`
}

// ListDomainsResult represents the ListDomainsResult schema from the OpenAPI specification
type ListDomainsResult struct {
	Domainnames interface{} `json:"DomainNames,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// NoSuchDomain represents the NoSuchDomain schema from the OpenAPI specification
type NoSuchDomain struct {
	Boxusage float32 `json:"BoxUsage,omitempty"`
}

// DeletableAttribute represents the DeletableAttribute schema from the OpenAPI specification
type DeletableAttribute struct {
	Value interface{} `json:"Value,omitempty"`
	Name interface{} `json:"Name"`
}

// DomainMetadataRequest represents the DomainMetadataRequest schema from the OpenAPI specification
type DomainMetadataRequest struct {
	Domainname interface{} `json:"DomainName"`
}

// InvalidNumberValueTests represents the InvalidNumberValueTests schema from the OpenAPI specification
type InvalidNumberValueTests struct {
	Boxusage float32 `json:"BoxUsage,omitempty"`
}

// NumberSubmittedAttributesExceeded represents the NumberSubmittedAttributesExceeded schema from the OpenAPI specification
type NumberSubmittedAttributesExceeded struct {
	Boxusage float32 `json:"BoxUsage,omitempty"`
}

// NumberDomainBytesExceeded represents the NumberDomainBytesExceeded schema from the OpenAPI specification
type NumberDomainBytesExceeded struct {
	Boxusage float32 `json:"BoxUsage,omitempty"`
}

// NumberSubmittedItemsExceeded represents the NumberSubmittedItemsExceeded schema from the OpenAPI specification
type NumberSubmittedItemsExceeded struct {
	Boxusage float32 `json:"BoxUsage,omitempty"`
}

// PutAttributesRequest represents the PutAttributesRequest schema from the OpenAPI specification
type PutAttributesRequest struct {
	Attributes interface{} `json:"Attributes"`
	Domainname interface{} `json:"DomainName"`
	Expected interface{} `json:"Expected,omitempty"`
	Itemname interface{} `json:"ItemName"`
}

// AttributeDoesNotExist represents the AttributeDoesNotExist schema from the OpenAPI specification
type AttributeDoesNotExist struct {
	Boxusage float32 `json:"BoxUsage,omitempty"`
}

// NumberDomainAttributesExceeded represents the NumberDomainAttributesExceeded schema from the OpenAPI specification
type NumberDomainAttributesExceeded struct {
	Boxusage float32 `json:"BoxUsage,omitempty"`
}

// DeleteAttributesRequest represents the DeleteAttributesRequest schema from the OpenAPI specification
type DeleteAttributesRequest struct {
	Itemname interface{} `json:"ItemName"`
	Attributes interface{} `json:"Attributes,omitempty"`
	Domainname interface{} `json:"DomainName"`
	Expected interface{} `json:"Expected,omitempty"`
}

// NumberItemAttributesExceeded represents the NumberItemAttributesExceeded schema from the OpenAPI specification
type NumberItemAttributesExceeded struct {
	Boxusage float32 `json:"BoxUsage,omitempty"`
}

// SelectRequest represents the SelectRequest schema from the OpenAPI specification
type SelectRequest struct {
	Selectexpression interface{} `json:"SelectExpression"`
	Consistentread interface{} `json:"ConsistentRead,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// MissingParameter represents the MissingParameter schema from the OpenAPI specification
type MissingParameter struct {
	Boxusage float32 `json:"BoxUsage,omitempty"`
}

// DeletableItem represents the DeletableItem schema from the OpenAPI specification
type DeletableItem struct {
	Name interface{} `json:"Name"`
	Attributes []interface{} `json:"Attributes,omitempty"`
}

// Item represents the Item schema from the OpenAPI specification
type Item struct {
	Name interface{} `json:"Name"`
	Alternatenameencoding interface{} `json:"AlternateNameEncoding,omitempty"`
	Attributes interface{} `json:"Attributes"`
}

// DuplicateItemName represents the DuplicateItemName schema from the OpenAPI specification
type DuplicateItemName struct {
	Boxusage float32 `json:"BoxUsage,omitempty"`
}

// InvalidNumberPredicates represents the InvalidNumberPredicates schema from the OpenAPI specification
type InvalidNumberPredicates struct {
	Boxusage float32 `json:"BoxUsage,omitempty"`
}

// DeleteDomainRequest represents the DeleteDomainRequest schema from the OpenAPI specification
type DeleteDomainRequest struct {
	Domainname interface{} `json:"DomainName"`
}

// ListDomainsRequest represents the ListDomainsRequest schema from the OpenAPI specification
type ListDomainsRequest struct {
	Maxnumberofdomains interface{} `json:"MaxNumberOfDomains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// InvalidParameterValue represents the InvalidParameterValue schema from the OpenAPI specification
type InvalidParameterValue struct {
	Boxusage float32 `json:"BoxUsage,omitempty"`
}

// BatchPutAttributesRequest represents the BatchPutAttributesRequest schema from the OpenAPI specification
type BatchPutAttributesRequest struct {
	Domainname interface{} `json:"DomainName"`
	Items interface{} `json:"Items"`
}

// InvalidNextToken represents the InvalidNextToken schema from the OpenAPI specification
type InvalidNextToken struct {
	Boxusage float32 `json:"BoxUsage,omitempty"`
}

// SelectResult represents the SelectResult schema from the OpenAPI specification
type SelectResult struct {
	Items interface{} `json:"Items,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DomainMetadataResult represents the DomainMetadataResult schema from the OpenAPI specification
type DomainMetadataResult struct {
	Itemcount interface{} `json:"ItemCount,omitempty"`
	Itemnamessizebytes interface{} `json:"ItemNamesSizeBytes,omitempty"`
	Timestamp interface{} `json:"Timestamp,omitempty"`
	Attributenamecount interface{} `json:"AttributeNameCount,omitempty"`
	Attributenamessizebytes interface{} `json:"AttributeNamesSizeBytes,omitempty"`
	Attributevaluecount interface{} `json:"AttributeValueCount,omitempty"`
	Attributevaluessizebytes interface{} `json:"AttributeValuesSizeBytes,omitempty"`
}

// CreateDomainRequest represents the CreateDomainRequest schema from the OpenAPI specification
type CreateDomainRequest struct {
	Domainname interface{} `json:"DomainName"`
}
