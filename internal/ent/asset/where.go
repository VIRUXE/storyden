// Code generated by ent, DO NOT EDIT.

package asset

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Southclaws/storyden/internal/ent/predicate"
	"github.com/rs/xid"
)

// ID filters vertices based on their ID field.
func ID(id xid.ID) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id xid.ID) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id xid.ID) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...xid.ID) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...xid.ID) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id xid.ID) predicate.Asset {
	return predicate.Asset(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id xid.ID) predicate.Asset {
	return predicate.Asset(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id xid.ID) predicate.Asset {
	return predicate.Asset(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id xid.ID) predicate.Asset {
	return predicate.Asset(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldUpdatedAt, v))
}

// Filename applies equality check predicate on the "filename" field. It's identical to FilenameEQ.
func Filename(v string) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldFilename, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldSize, v))
}

// AccountID applies equality check predicate on the "account_id" field. It's identical to AccountIDEQ.
func AccountID(v xid.ID) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldAccountID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldLTE(FieldUpdatedAt, v))
}

// FilenameEQ applies the EQ predicate on the "filename" field.
func FilenameEQ(v string) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldFilename, v))
}

// FilenameNEQ applies the NEQ predicate on the "filename" field.
func FilenameNEQ(v string) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldFilename, v))
}

// FilenameIn applies the In predicate on the "filename" field.
func FilenameIn(vs ...string) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldFilename, vs...))
}

// FilenameNotIn applies the NotIn predicate on the "filename" field.
func FilenameNotIn(vs ...string) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldFilename, vs...))
}

// FilenameGT applies the GT predicate on the "filename" field.
func FilenameGT(v string) predicate.Asset {
	return predicate.Asset(sql.FieldGT(FieldFilename, v))
}

// FilenameGTE applies the GTE predicate on the "filename" field.
func FilenameGTE(v string) predicate.Asset {
	return predicate.Asset(sql.FieldGTE(FieldFilename, v))
}

// FilenameLT applies the LT predicate on the "filename" field.
func FilenameLT(v string) predicate.Asset {
	return predicate.Asset(sql.FieldLT(FieldFilename, v))
}

// FilenameLTE applies the LTE predicate on the "filename" field.
func FilenameLTE(v string) predicate.Asset {
	return predicate.Asset(sql.FieldLTE(FieldFilename, v))
}

// FilenameContains applies the Contains predicate on the "filename" field.
func FilenameContains(v string) predicate.Asset {
	return predicate.Asset(sql.FieldContains(FieldFilename, v))
}

// FilenameHasPrefix applies the HasPrefix predicate on the "filename" field.
func FilenameHasPrefix(v string) predicate.Asset {
	return predicate.Asset(sql.FieldHasPrefix(FieldFilename, v))
}

// FilenameHasSuffix applies the HasSuffix predicate on the "filename" field.
func FilenameHasSuffix(v string) predicate.Asset {
	return predicate.Asset(sql.FieldHasSuffix(FieldFilename, v))
}

// FilenameEqualFold applies the EqualFold predicate on the "filename" field.
func FilenameEqualFold(v string) predicate.Asset {
	return predicate.Asset(sql.FieldEqualFold(FieldFilename, v))
}

// FilenameContainsFold applies the ContainsFold predicate on the "filename" field.
func FilenameContainsFold(v string) predicate.Asset {
	return predicate.Asset(sql.FieldContainsFold(FieldFilename, v))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int) predicate.Asset {
	return predicate.Asset(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int) predicate.Asset {
	return predicate.Asset(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int) predicate.Asset {
	return predicate.Asset(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int) predicate.Asset {
	return predicate.Asset(sql.FieldLTE(FieldSize, v))
}

// MetadataIsNil applies the IsNil predicate on the "metadata" field.
func MetadataIsNil() predicate.Asset {
	return predicate.Asset(sql.FieldIsNull(FieldMetadata))
}

// MetadataNotNil applies the NotNil predicate on the "metadata" field.
func MetadataNotNil() predicate.Asset {
	return predicate.Asset(sql.FieldNotNull(FieldMetadata))
}

// AccountIDEQ applies the EQ predicate on the "account_id" field.
func AccountIDEQ(v xid.ID) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldAccountID, v))
}

// AccountIDNEQ applies the NEQ predicate on the "account_id" field.
func AccountIDNEQ(v xid.ID) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldAccountID, v))
}

// AccountIDIn applies the In predicate on the "account_id" field.
func AccountIDIn(vs ...xid.ID) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldAccountID, vs...))
}

// AccountIDNotIn applies the NotIn predicate on the "account_id" field.
func AccountIDNotIn(vs ...xid.ID) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldAccountID, vs...))
}

// AccountIDGT applies the GT predicate on the "account_id" field.
func AccountIDGT(v xid.ID) predicate.Asset {
	return predicate.Asset(sql.FieldGT(FieldAccountID, v))
}

// AccountIDGTE applies the GTE predicate on the "account_id" field.
func AccountIDGTE(v xid.ID) predicate.Asset {
	return predicate.Asset(sql.FieldGTE(FieldAccountID, v))
}

// AccountIDLT applies the LT predicate on the "account_id" field.
func AccountIDLT(v xid.ID) predicate.Asset {
	return predicate.Asset(sql.FieldLT(FieldAccountID, v))
}

// AccountIDLTE applies the LTE predicate on the "account_id" field.
func AccountIDLTE(v xid.ID) predicate.Asset {
	return predicate.Asset(sql.FieldLTE(FieldAccountID, v))
}

// AccountIDContains applies the Contains predicate on the "account_id" field.
func AccountIDContains(v xid.ID) predicate.Asset {
	vc := v.String()
	return predicate.Asset(sql.FieldContains(FieldAccountID, vc))
}

// AccountIDHasPrefix applies the HasPrefix predicate on the "account_id" field.
func AccountIDHasPrefix(v xid.ID) predicate.Asset {
	vc := v.String()
	return predicate.Asset(sql.FieldHasPrefix(FieldAccountID, vc))
}

// AccountIDHasSuffix applies the HasSuffix predicate on the "account_id" field.
func AccountIDHasSuffix(v xid.ID) predicate.Asset {
	vc := v.String()
	return predicate.Asset(sql.FieldHasSuffix(FieldAccountID, vc))
}

// AccountIDEqualFold applies the EqualFold predicate on the "account_id" field.
func AccountIDEqualFold(v xid.ID) predicate.Asset {
	vc := v.String()
	return predicate.Asset(sql.FieldEqualFold(FieldAccountID, vc))
}

// AccountIDContainsFold applies the ContainsFold predicate on the "account_id" field.
func AccountIDContainsFold(v xid.ID) predicate.Asset {
	vc := v.String()
	return predicate.Asset(sql.FieldContainsFold(FieldAccountID, vc))
}

// HasPosts applies the HasEdge predicate on the "posts" edge.
func HasPosts() predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PostsTable, PostsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostsWith applies the HasEdge predicate on the "posts" edge with a given conditions (other predicates).
func HasPostsWith(preds ...predicate.Post) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := newPostsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNodes applies the HasEdge predicate on the "nodes" edge.
func HasNodes() predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, NodesTable, NodesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNodesWith applies the HasEdge predicate on the "nodes" edge with a given conditions (other predicates).
func HasNodesWith(preds ...predicate.Node) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := newNodesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLinks applies the HasEdge predicate on the "links" edge.
func HasLinks() predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, LinksTable, LinksPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLinksWith applies the HasEdge predicate on the "links" edge with a given conditions (other predicates).
func HasLinksWith(preds ...predicate.Link) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := newLinksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Account) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEvent applies the HasEdge predicate on the "event" edge.
func HasEvent() predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EventTable, EventColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventWith applies the HasEdge predicate on the "event" edge with a given conditions (other predicates).
func HasEventWith(preds ...predicate.Event) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := newEventStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Asset) predicate.Asset {
	return predicate.Asset(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Asset) predicate.Asset {
	return predicate.Asset(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Asset) predicate.Asset {
	return predicate.Asset(sql.NotPredicates(p))
}
