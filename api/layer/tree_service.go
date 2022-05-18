package layer

import (
	"context"
	"errors"

	"github.com/nspcc-dev/neofs-s3-gw/api/data"
	cid "github.com/nspcc-dev/neofs-sdk-go/container/id"
	oid "github.com/nspcc-dev/neofs-sdk-go/object/id"
)

// TreeService provide interface to interact with tree service using s3 data models.
type TreeService interface {
	// PutSettingsNode update or create new settings node in tree service.
	PutSettingsNode(context.Context, *cid.ID, *data.BucketSettings) error

	// GetSettingsNode retrieves the settings node from the tree service and form data.BucketSettings.
	//
	// If node is not found returns ErrNodeNotFound error.
	GetSettingsNode(context.Context, *cid.ID) (*data.BucketSettings, error)

	GetNotificationConfigurationNode(ctx context.Context, cnrID *cid.ID) (*oid.ID, error)
	// PutNotificationConfigurationNode puts a node to a system tree
	// and returns objectID of a previous notif config which must be deleted in NeoFS
	PutNotificationConfigurationNode(ctx context.Context, cnrID *cid.ID, objID *oid.ID) (*oid.ID, error)

	GetBucketCORS(ctx context.Context, cnrID *cid.ID) (*oid.ID, error)
	// PutBucketCORS puts a node to a system tree and returns objectID of a previous cors config which must be deleted in NeoFS
	PutBucketCORS(ctx context.Context, cnrID *cid.ID, objID *oid.ID) (*oid.ID, error)
	// DeleteBucketCORS removes a node from a system tree and returns objID which must be deleted in NeoFS
	DeleteBucketCORS(ctx context.Context, cnrID *cid.ID) (*oid.ID, error)

	GetVersions(ctx context.Context, cnrID *cid.ID, objectName string) ([]*NodeVersion, error)
	GetLatestVersion(ctx context.Context, cnrID *cid.ID, objectName string) (*NodeVersion, error)
	GetUnversioned(ctx context.Context, cnrID *cid.ID, objectName string) (*NodeVersion, error)
	AddVersion(ctx context.Context, cnrID *cid.ID, objectName string, newVersion *NodeVersion) error
	RemoveVersion(ctx context.Context, cnrID *cid.ID, nodeID uint64) error

	AddSystemVersion(ctx context.Context, cnrID *cid.ID, objectName string, newVersion *BaseNodeVersion) error
	GetSystemVersion(ctx context.Context, cnrID *cid.ID, objectName string) (*BaseNodeVersion, error)
	RemoveSystemVersion(ctx context.Context, cnrID *cid.ID, nodeID uint64) error
}

type NodeVersion struct {
	BaseNodeVersion
	IsDeleteMarker bool
	IsUnversioned  bool
}

type BaseNodeVersion struct {
	ID  uint64
	OID oid.ID
}

// ErrNodeNotFound is returned from Tree service in case of not found error.
var ErrNodeNotFound = errors.New("not found")