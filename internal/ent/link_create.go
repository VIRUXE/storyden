// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/asset"
	"github.com/Southclaws/storyden/internal/ent/cluster"
	"github.com/Southclaws/storyden/internal/ent/item"
	"github.com/Southclaws/storyden/internal/ent/link"
	"github.com/Southclaws/storyden/internal/ent/post"
	"github.com/rs/xid"
)

// LinkCreate is the builder for creating a Link entity.
type LinkCreate struct {
	config
	mutation *LinkMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (lc *LinkCreate) SetCreatedAt(t time.Time) *LinkCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *LinkCreate) SetNillableCreatedAt(t *time.Time) *LinkCreate {
	if t != nil {
		lc.SetCreatedAt(*t)
	}
	return lc
}

// SetURL sets the "url" field.
func (lc *LinkCreate) SetURL(s string) *LinkCreate {
	lc.mutation.SetURL(s)
	return lc
}

// SetTitle sets the "title" field.
func (lc *LinkCreate) SetTitle(s string) *LinkCreate {
	lc.mutation.SetTitle(s)
	return lc
}

// SetDescription sets the "description" field.
func (lc *LinkCreate) SetDescription(s string) *LinkCreate {
	lc.mutation.SetDescription(s)
	return lc
}

// SetID sets the "id" field.
func (lc *LinkCreate) SetID(x xid.ID) *LinkCreate {
	lc.mutation.SetID(x)
	return lc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (lc *LinkCreate) SetNillableID(x *xid.ID) *LinkCreate {
	if x != nil {
		lc.SetID(*x)
	}
	return lc
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (lc *LinkCreate) AddPostIDs(ids ...xid.ID) *LinkCreate {
	lc.mutation.AddPostIDs(ids...)
	return lc
}

// AddPosts adds the "posts" edges to the Post entity.
func (lc *LinkCreate) AddPosts(p ...*Post) *LinkCreate {
	ids := make([]xid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return lc.AddPostIDs(ids...)
}

// AddClusterIDs adds the "clusters" edge to the Cluster entity by IDs.
func (lc *LinkCreate) AddClusterIDs(ids ...xid.ID) *LinkCreate {
	lc.mutation.AddClusterIDs(ids...)
	return lc
}

// AddClusters adds the "clusters" edges to the Cluster entity.
func (lc *LinkCreate) AddClusters(c ...*Cluster) *LinkCreate {
	ids := make([]xid.ID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return lc.AddClusterIDs(ids...)
}

// AddItemIDs adds the "items" edge to the Item entity by IDs.
func (lc *LinkCreate) AddItemIDs(ids ...xid.ID) *LinkCreate {
	lc.mutation.AddItemIDs(ids...)
	return lc
}

// AddItems adds the "items" edges to the Item entity.
func (lc *LinkCreate) AddItems(i ...*Item) *LinkCreate {
	ids := make([]xid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return lc.AddItemIDs(ids...)
}

// AddAssetIDs adds the "assets" edge to the Asset entity by IDs.
func (lc *LinkCreate) AddAssetIDs(ids ...string) *LinkCreate {
	lc.mutation.AddAssetIDs(ids...)
	return lc
}

// AddAssets adds the "assets" edges to the Asset entity.
func (lc *LinkCreate) AddAssets(a ...*Asset) *LinkCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return lc.AddAssetIDs(ids...)
}

// Mutation returns the LinkMutation object of the builder.
func (lc *LinkCreate) Mutation() *LinkMutation {
	return lc.mutation
}

// Save creates the Link in the database.
func (lc *LinkCreate) Save(ctx context.Context) (*Link, error) {
	lc.defaults()
	return withHooks(ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LinkCreate) SaveX(ctx context.Context) *Link {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LinkCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LinkCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LinkCreate) defaults() {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		v := link.DefaultCreatedAt()
		lc.mutation.SetCreatedAt(v)
	}
	if _, ok := lc.mutation.ID(); !ok {
		v := link.DefaultID()
		lc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LinkCreate) check() error {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Link.created_at"`)}
	}
	if _, ok := lc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Link.url"`)}
	}
	if _, ok := lc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Link.title"`)}
	}
	if _, ok := lc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Link.description"`)}
	}
	if v, ok := lc.mutation.ID(); ok {
		if err := link.IDValidator(v.String()); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Link.id": %w`, err)}
		}
	}
	return nil
}

func (lc *LinkCreate) sqlSave(ctx context.Context) (*Link, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*xid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LinkCreate) createSpec() (*Link, *sqlgraph.CreateSpec) {
	var (
		_node = &Link{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(link.Table, sqlgraph.NewFieldSpec(link.FieldID, field.TypeString))
	)
	_spec.OnConflict = lc.conflict
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.SetField(link.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lc.mutation.URL(); ok {
		_spec.SetField(link.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := lc.mutation.Title(); ok {
		_spec.SetField(link.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := lc.mutation.Description(); ok {
		_spec.SetField(link.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := lc.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   link.PostsTable,
			Columns: link.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.ClustersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   link.ClustersTable,
			Columns: link.ClustersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cluster.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   link.ItemsTable,
			Columns: link.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.AssetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   link.AssetsTable,
			Columns: link.AssetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Link.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LinkUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (lc *LinkCreate) OnConflict(opts ...sql.ConflictOption) *LinkUpsertOne {
	lc.conflict = opts
	return &LinkUpsertOne{
		create: lc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Link.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lc *LinkCreate) OnConflictColumns(columns ...string) *LinkUpsertOne {
	lc.conflict = append(lc.conflict, sql.ConflictColumns(columns...))
	return &LinkUpsertOne{
		create: lc,
	}
}

type (
	// LinkUpsertOne is the builder for "upsert"-ing
	//  one Link node.
	LinkUpsertOne struct {
		create *LinkCreate
	}

	// LinkUpsert is the "OnConflict" setter.
	LinkUpsert struct {
		*sql.UpdateSet
	}
)

// SetTitle sets the "title" field.
func (u *LinkUpsert) SetTitle(v string) *LinkUpsert {
	u.Set(link.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *LinkUpsert) UpdateTitle() *LinkUpsert {
	u.SetExcluded(link.FieldTitle)
	return u
}

// SetDescription sets the "description" field.
func (u *LinkUpsert) SetDescription(v string) *LinkUpsert {
	u.Set(link.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *LinkUpsert) UpdateDescription() *LinkUpsert {
	u.SetExcluded(link.FieldDescription)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Link.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(link.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *LinkUpsertOne) UpdateNewValues() *LinkUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(link.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(link.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.URL(); exists {
			s.SetIgnore(link.FieldURL)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Link.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *LinkUpsertOne) Ignore() *LinkUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LinkUpsertOne) DoNothing() *LinkUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LinkCreate.OnConflict
// documentation for more info.
func (u *LinkUpsertOne) Update(set func(*LinkUpsert)) *LinkUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LinkUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *LinkUpsertOne) SetTitle(v string) *LinkUpsertOne {
	return u.Update(func(s *LinkUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *LinkUpsertOne) UpdateTitle() *LinkUpsertOne {
	return u.Update(func(s *LinkUpsert) {
		s.UpdateTitle()
	})
}

// SetDescription sets the "description" field.
func (u *LinkUpsertOne) SetDescription(v string) *LinkUpsertOne {
	return u.Update(func(s *LinkUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *LinkUpsertOne) UpdateDescription() *LinkUpsertOne {
	return u.Update(func(s *LinkUpsert) {
		s.UpdateDescription()
	})
}

// Exec executes the query.
func (u *LinkUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LinkCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LinkUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *LinkUpsertOne) ID(ctx context.Context) (id xid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: LinkUpsertOne.ID is not supported by MySQL driver. Use LinkUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *LinkUpsertOne) IDX(ctx context.Context) xid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// LinkCreateBulk is the builder for creating many Link entities in bulk.
type LinkCreateBulk struct {
	config
	err      error
	builders []*LinkCreate
	conflict []sql.ConflictOption
}

// Save creates the Link entities in the database.
func (lcb *LinkCreateBulk) Save(ctx context.Context) ([]*Link, error) {
	if lcb.err != nil {
		return nil, lcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Link, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LinkMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = lcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LinkCreateBulk) SaveX(ctx context.Context) []*Link {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LinkCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LinkCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Link.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LinkUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (lcb *LinkCreateBulk) OnConflict(opts ...sql.ConflictOption) *LinkUpsertBulk {
	lcb.conflict = opts
	return &LinkUpsertBulk{
		create: lcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Link.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lcb *LinkCreateBulk) OnConflictColumns(columns ...string) *LinkUpsertBulk {
	lcb.conflict = append(lcb.conflict, sql.ConflictColumns(columns...))
	return &LinkUpsertBulk{
		create: lcb,
	}
}

// LinkUpsertBulk is the builder for "upsert"-ing
// a bulk of Link nodes.
type LinkUpsertBulk struct {
	create *LinkCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Link.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(link.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *LinkUpsertBulk) UpdateNewValues() *LinkUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(link.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(link.FieldCreatedAt)
			}
			if _, exists := b.mutation.URL(); exists {
				s.SetIgnore(link.FieldURL)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Link.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *LinkUpsertBulk) Ignore() *LinkUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LinkUpsertBulk) DoNothing() *LinkUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LinkCreateBulk.OnConflict
// documentation for more info.
func (u *LinkUpsertBulk) Update(set func(*LinkUpsert)) *LinkUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LinkUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *LinkUpsertBulk) SetTitle(v string) *LinkUpsertBulk {
	return u.Update(func(s *LinkUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *LinkUpsertBulk) UpdateTitle() *LinkUpsertBulk {
	return u.Update(func(s *LinkUpsert) {
		s.UpdateTitle()
	})
}

// SetDescription sets the "description" field.
func (u *LinkUpsertBulk) SetDescription(v string) *LinkUpsertBulk {
	return u.Update(func(s *LinkUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *LinkUpsertBulk) UpdateDescription() *LinkUpsertBulk {
	return u.Update(func(s *LinkUpsert) {
		s.UpdateDescription()
	})
}

// Exec executes the query.
func (u *LinkUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the LinkCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LinkCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LinkUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
