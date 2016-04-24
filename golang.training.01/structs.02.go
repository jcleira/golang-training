// Representation of a message.
import (
	"github.com/getsentry/raven-go"
	"github.com/keradgames/kg-aws-golang-client/dynamo"
)

type Message struct {
	ID           string      `json:"id" dynamodb:"id"`
	AccountID    string      `json:"-" dynamodb:"account_id"`
	DisplayType  string      `json:"display_type" dynamodb:"display_type"`
	Category     Category    `json:"category" dynamodb:"category"`
	Read         bool        `json:"read" dynamodb:"has_read"`
	ExpiredAt    int64       `json:"-" dynamodb:"expired_at"`
	CreatedAt    int64       `json:"created_at" dynamodb:"created_at"`
	UpdatedAt    int64       `json:"updated_at" dynamodb:"updated_at"`
	Metadata     interface{} `json:"metadata,omitempty" dynamodb:"-"`
	MetadataJSON string      `json:"-" dynamodb:"metadata"`
}

type Client struct {
	DynamoAccessor dynamo.DynamoAccessor
	DynamoTables   map[string]string
	Raven          *raven.Client
}
