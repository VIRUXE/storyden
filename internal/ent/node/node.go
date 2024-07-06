// Code generated by ent, DO NOT EDIT.

package node

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/rs/xid"
)

const (
	// Label holds the string label denoting the node type in the database.
	Label = "node"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldParentNodeID holds the string denoting the parent_node_id field in the database.
	FieldParentNodeID = "parent_node_id"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldVisibility holds the string denoting the visibility field in the database.
	FieldVisibility = "visibility"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeNodes holds the string denoting the nodes edge name in mutations.
	EdgeNodes = "nodes"
	// EdgeAssets holds the string denoting the assets edge name in mutations.
	EdgeAssets = "assets"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeLinks holds the string denoting the links edge name in mutations.
	EdgeLinks = "links"
	// EdgeCollections holds the string denoting the collections edge name in mutations.
	EdgeCollections = "collections"
	// Table holds the table name of the node in the database.
	Table = "nodes"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "nodes"
	// OwnerInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	OwnerInverseTable = "accounts"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "account_id"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "nodes"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_node_id"
	// NodesTable is the table that holds the nodes relation/edge.
	NodesTable = "nodes"
	// NodesColumn is the table column denoting the nodes relation/edge.
	NodesColumn = "parent_node_id"
	// AssetsTable is the table that holds the assets relation/edge. The primary key declared below.
	AssetsTable = "node_assets"
	// AssetsInverseTable is the table name for the Asset entity.
	// It exists in this package in order to avoid circular dependency with the "asset" package.
	AssetsInverseTable = "assets"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "tag_nodes"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// LinksTable is the table that holds the links relation/edge. The primary key declared below.
	LinksTable = "link_nodes"
	// LinksInverseTable is the table name for the Link entity.
	// It exists in this package in order to avoid circular dependency with the "link" package.
	LinksInverseTable = "links"
	// CollectionsTable is the table that holds the collections relation/edge. The primary key declared below.
	CollectionsTable = "collection_nodes"
	// CollectionsInverseTable is the table name for the Collection entity.
	// It exists in this package in order to avoid circular dependency with the "collection" package.
	CollectionsInverseTable = "collections"
)

// Columns holds all SQL columns for node fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldSlug,
	FieldDescription,
	FieldContent,
	FieldParentNodeID,
	FieldAccountID,
	FieldVisibility,
	FieldMetadata,
}

var (
	// AssetsPrimaryKey and AssetsColumn2 are the table columns denoting the
	// primary key for the assets relation (M2M).
	AssetsPrimaryKey = []string{"node_id", "asset_id"}
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"tag_id", "node_id"}
	// LinksPrimaryKey and LinksColumn2 are the table columns denoting the
	// primary key for the links relation (M2M).
	LinksPrimaryKey = []string{"link_id", "node_id"}
	// CollectionsPrimaryKey and CollectionsColumn2 are the table columns denoting the
	// primary key for the collections relation (M2M).
	CollectionsPrimaryKey = []string{"collection_id", "node_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() xid.ID
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// Visibility defines the type for the "visibility" enum field.
type Visibility string

// VisibilityDraft is the default value of the Visibility enum.
const DefaultVisibility = VisibilityDraft

// Visibility values.
const (
	VisibilityDraft     Visibility = "draft"
	VisibilityUnlisted  Visibility = "unlisted"
	VisibilityReview    Visibility = "review"
	VisibilityPublished Visibility = "published"
)

func (v Visibility) String() string {
	return string(v)
}

// VisibilityValidator is a validator for the "visibility" field enum values. It is called by the builders before save.
func VisibilityValidator(v Visibility) error {
	switch v {
	case VisibilityDraft, VisibilityUnlisted, VisibilityReview, VisibilityPublished:
		return nil
	default:
		return fmt.Errorf("node: invalid enum value for visibility field: %q", v)
	}
}

// OrderOption defines the ordering options for the Node queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySlug orders the results by the slug field.
func BySlug(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSlug, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByParentNodeID orders the results by the parent_node_id field.
func ByParentNodeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentNodeID, opts...).ToFunc()
}

// ByAccountID orders the results by the account_id field.
func ByAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountID, opts...).ToFunc()
}

// ByVisibility orders the results by the visibility field.
func ByVisibility(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVisibility, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByParentField orders the results by parent field.
func ByParentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), sql.OrderByField(field, opts...))
	}
}

// ByNodesCount orders the results by nodes count.
func ByNodesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNodesStep(), opts...)
	}
}

// ByNodes orders the results by nodes terms.
func ByNodes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNodesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAssetsCount orders the results by assets count.
func ByAssetsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAssetsStep(), opts...)
	}
}

// ByAssets orders the results by assets terms.
func ByAssets(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAssetsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTagsCount orders the results by tags count.
func ByTagsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTagsStep(), opts...)
	}
}

// ByTags orders the results by tags terms.
func ByTags(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLinksCount orders the results by links count.
func ByLinksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLinksStep(), opts...)
	}
}

// ByLinks orders the results by links terms.
func ByLinks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLinksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCollectionsCount orders the results by collections count.
func ByCollectionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCollectionsStep(), opts...)
	}
}

// ByCollections orders the results by collections terms.
func ByCollections(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCollectionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
func newParentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
	)
}
func newNodesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NodesTable, NodesColumn),
	)
}
func newAssetsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AssetsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, AssetsTable, AssetsPrimaryKey...),
	)
}
func newTagsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, TagsTable, TagsPrimaryKey...),
	)
}
func newLinksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LinksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, LinksTable, LinksPrimaryKey...),
	)
}
func newCollectionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CollectionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, CollectionsTable, CollectionsPrimaryKey...),
	)
}
