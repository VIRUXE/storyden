// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/asset"
	"github.com/Southclaws/storyden/internal/ent/link"
	"github.com/Southclaws/storyden/internal/ent/node"
	"github.com/Southclaws/storyden/internal/ent/post"
	"github.com/Southclaws/storyden/internal/ent/predicate"
	"github.com/rs/xid"
)

// LinkQuery is the builder for querying Link entities.
type LinkQuery struct {
	config
	ctx                       *QueryContext
	order                     []link.OrderOption
	inters                    []Interceptor
	predicates                []predicate.Link
	withPosts                 *PostQuery
	withPostContentReferences *PostQuery
	withNodes                 *NodeQuery
	withNodeContentReferences *NodeQuery
	withAssets                *AssetQuery
	modifiers                 []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LinkQuery builder.
func (lq *LinkQuery) Where(ps ...predicate.Link) *LinkQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit the number of records to be returned by this query.
func (lq *LinkQuery) Limit(limit int) *LinkQuery {
	lq.ctx.Limit = &limit
	return lq
}

// Offset to start from.
func (lq *LinkQuery) Offset(offset int) *LinkQuery {
	lq.ctx.Offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *LinkQuery) Unique(unique bool) *LinkQuery {
	lq.ctx.Unique = &unique
	return lq
}

// Order specifies how the records should be ordered.
func (lq *LinkQuery) Order(o ...link.OrderOption) *LinkQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// QueryPosts chains the current query on the "posts" edge.
func (lq *LinkQuery) QueryPosts() *PostQuery {
	query := (&PostClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(link.Table, link.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, link.PostsTable, link.PostsColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPostContentReferences chains the current query on the "post_content_references" edge.
func (lq *LinkQuery) QueryPostContentReferences() *PostQuery {
	query := (&PostClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(link.Table, link.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, link.PostContentReferencesTable, link.PostContentReferencesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNodes chains the current query on the "nodes" edge.
func (lq *LinkQuery) QueryNodes() *NodeQuery {
	query := (&NodeClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(link.Table, link.FieldID, selector),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, link.NodesTable, link.NodesColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNodeContentReferences chains the current query on the "node_content_references" edge.
func (lq *LinkQuery) QueryNodeContentReferences() *NodeQuery {
	query := (&NodeClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(link.Table, link.FieldID, selector),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, link.NodeContentReferencesTable, link.NodeContentReferencesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAssets chains the current query on the "assets" edge.
func (lq *LinkQuery) QueryAssets() *AssetQuery {
	query := (&AssetClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(link.Table, link.FieldID, selector),
			sqlgraph.To(asset.Table, asset.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, link.AssetsTable, link.AssetsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Link entity from the query.
// Returns a *NotFoundError when no Link was found.
func (lq *LinkQuery) First(ctx context.Context) (*Link, error) {
	nodes, err := lq.Limit(1).All(setContextOp(ctx, lq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{link.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LinkQuery) FirstX(ctx context.Context) *Link {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Link ID from the query.
// Returns a *NotFoundError when no Link ID was found.
func (lq *LinkQuery) FirstID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = lq.Limit(1).IDs(setContextOp(ctx, lq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{link.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lq *LinkQuery) FirstIDX(ctx context.Context) xid.ID {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Link entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Link entity is found.
// Returns a *NotFoundError when no Link entities are found.
func (lq *LinkQuery) Only(ctx context.Context) (*Link, error) {
	nodes, err := lq.Limit(2).All(setContextOp(ctx, lq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{link.Label}
	default:
		return nil, &NotSingularError{link.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LinkQuery) OnlyX(ctx context.Context) *Link {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Link ID in the query.
// Returns a *NotSingularError when more than one Link ID is found.
// Returns a *NotFoundError when no entities are found.
func (lq *LinkQuery) OnlyID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = lq.Limit(2).IDs(setContextOp(ctx, lq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{link.Label}
	default:
		err = &NotSingularError{link.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *LinkQuery) OnlyIDX(ctx context.Context) xid.ID {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Links.
func (lq *LinkQuery) All(ctx context.Context) ([]*Link, error) {
	ctx = setContextOp(ctx, lq.ctx, "All")
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Link, *LinkQuery]()
	return withInterceptors[[]*Link](ctx, lq, qr, lq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lq *LinkQuery) AllX(ctx context.Context) []*Link {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Link IDs.
func (lq *LinkQuery) IDs(ctx context.Context) (ids []xid.ID, err error) {
	if lq.ctx.Unique == nil && lq.path != nil {
		lq.Unique(true)
	}
	ctx = setContextOp(ctx, lq.ctx, "IDs")
	if err = lq.Select(link.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *LinkQuery) IDsX(ctx context.Context) []xid.ID {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *LinkQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lq.ctx, "Count")
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lq, querierCount[*LinkQuery](), lq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LinkQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LinkQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lq.ctx, "Exist")
	switch _, err := lq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LinkQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LinkQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LinkQuery) Clone() *LinkQuery {
	if lq == nil {
		return nil
	}
	return &LinkQuery{
		config:                    lq.config,
		ctx:                       lq.ctx.Clone(),
		order:                     append([]link.OrderOption{}, lq.order...),
		inters:                    append([]Interceptor{}, lq.inters...),
		predicates:                append([]predicate.Link{}, lq.predicates...),
		withPosts:                 lq.withPosts.Clone(),
		withPostContentReferences: lq.withPostContentReferences.Clone(),
		withNodes:                 lq.withNodes.Clone(),
		withNodeContentReferences: lq.withNodeContentReferences.Clone(),
		withAssets:                lq.withAssets.Clone(),
		// clone intermediate query.
		sql:  lq.sql.Clone(),
		path: lq.path,
	}
}

// WithPosts tells the query-builder to eager-load the nodes that are connected to
// the "posts" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LinkQuery) WithPosts(opts ...func(*PostQuery)) *LinkQuery {
	query := (&PostClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withPosts = query
	return lq
}

// WithPostContentReferences tells the query-builder to eager-load the nodes that are connected to
// the "post_content_references" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LinkQuery) WithPostContentReferences(opts ...func(*PostQuery)) *LinkQuery {
	query := (&PostClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withPostContentReferences = query
	return lq
}

// WithNodes tells the query-builder to eager-load the nodes that are connected to
// the "nodes" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LinkQuery) WithNodes(opts ...func(*NodeQuery)) *LinkQuery {
	query := (&NodeClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withNodes = query
	return lq
}

// WithNodeContentReferences tells the query-builder to eager-load the nodes that are connected to
// the "node_content_references" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LinkQuery) WithNodeContentReferences(opts ...func(*NodeQuery)) *LinkQuery {
	query := (&NodeClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withNodeContentReferences = query
	return lq
}

// WithAssets tells the query-builder to eager-load the nodes that are connected to
// the "assets" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LinkQuery) WithAssets(opts ...func(*AssetQuery)) *LinkQuery {
	query := (&AssetClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withAssets = query
	return lq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Link.Query().
//		GroupBy(link.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lq *LinkQuery) GroupBy(field string, fields ...string) *LinkGroupBy {
	lq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LinkGroupBy{build: lq}
	grbuild.flds = &lq.ctx.Fields
	grbuild.label = link.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Link.Query().
//		Select(link.FieldCreatedAt).
//		Scan(ctx, &v)
func (lq *LinkQuery) Select(fields ...string) *LinkSelect {
	lq.ctx.Fields = append(lq.ctx.Fields, fields...)
	sbuild := &LinkSelect{LinkQuery: lq}
	sbuild.label = link.Label
	sbuild.flds, sbuild.scan = &lq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LinkSelect configured with the given aggregations.
func (lq *LinkQuery) Aggregate(fns ...AggregateFunc) *LinkSelect {
	return lq.Select().Aggregate(fns...)
}

func (lq *LinkQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lq); err != nil {
				return err
			}
		}
	}
	for _, f := range lq.ctx.Fields {
		if !link.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *LinkQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Link, error) {
	var (
		nodes       = []*Link{}
		_spec       = lq.querySpec()
		loadedTypes = [5]bool{
			lq.withPosts != nil,
			lq.withPostContentReferences != nil,
			lq.withNodes != nil,
			lq.withNodeContentReferences != nil,
			lq.withAssets != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Link).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Link{config: lq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(lq.modifiers) > 0 {
		_spec.Modifiers = lq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lq.withPosts; query != nil {
		if err := lq.loadPosts(ctx, query, nodes,
			func(n *Link) { n.Edges.Posts = []*Post{} },
			func(n *Link, e *Post) { n.Edges.Posts = append(n.Edges.Posts, e) }); err != nil {
			return nil, err
		}
	}
	if query := lq.withPostContentReferences; query != nil {
		if err := lq.loadPostContentReferences(ctx, query, nodes,
			func(n *Link) { n.Edges.PostContentReferences = []*Post{} },
			func(n *Link, e *Post) { n.Edges.PostContentReferences = append(n.Edges.PostContentReferences, e) }); err != nil {
			return nil, err
		}
	}
	if query := lq.withNodes; query != nil {
		if err := lq.loadNodes(ctx, query, nodes,
			func(n *Link) { n.Edges.Nodes = []*Node{} },
			func(n *Link, e *Node) { n.Edges.Nodes = append(n.Edges.Nodes, e) }); err != nil {
			return nil, err
		}
	}
	if query := lq.withNodeContentReferences; query != nil {
		if err := lq.loadNodeContentReferences(ctx, query, nodes,
			func(n *Link) { n.Edges.NodeContentReferences = []*Node{} },
			func(n *Link, e *Node) { n.Edges.NodeContentReferences = append(n.Edges.NodeContentReferences, e) }); err != nil {
			return nil, err
		}
	}
	if query := lq.withAssets; query != nil {
		if err := lq.loadAssets(ctx, query, nodes,
			func(n *Link) { n.Edges.Assets = []*Asset{} },
			func(n *Link, e *Asset) { n.Edges.Assets = append(n.Edges.Assets, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lq *LinkQuery) loadPosts(ctx context.Context, query *PostQuery, nodes []*Link, init func(*Link), assign func(*Link, *Post)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[xid.ID]*Link)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(post.FieldLinkID)
	}
	query.Where(predicate.Post(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(link.PostsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.LinkID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "link_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (lq *LinkQuery) loadPostContentReferences(ctx context.Context, query *PostQuery, nodes []*Link, init func(*Link), assign func(*Link, *Post)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[xid.ID]*Link)
	nids := make(map[xid.ID]map[*Link]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(link.PostContentReferencesTable)
		s.Join(joinT).On(s.C(post.FieldID), joinT.C(link.PostContentReferencesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(link.PostContentReferencesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(link.PostContentReferencesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(xid.ID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*xid.ID)
				inValue := *values[1].(*xid.ID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Link]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Post](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "post_content_references" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (lq *LinkQuery) loadNodes(ctx context.Context, query *NodeQuery, nodes []*Link, init func(*Link), assign func(*Link, *Node)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[xid.ID]*Link)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(node.FieldLinkID)
	}
	query.Where(predicate.Node(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(link.NodesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.LinkID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "link_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (lq *LinkQuery) loadNodeContentReferences(ctx context.Context, query *NodeQuery, nodes []*Link, init func(*Link), assign func(*Link, *Node)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[xid.ID]*Link)
	nids := make(map[xid.ID]map[*Link]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(link.NodeContentReferencesTable)
		s.Join(joinT).On(s.C(node.FieldID), joinT.C(link.NodeContentReferencesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(link.NodeContentReferencesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(link.NodeContentReferencesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(xid.ID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*xid.ID)
				inValue := *values[1].(*xid.ID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Link]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Node](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "node_content_references" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (lq *LinkQuery) loadAssets(ctx context.Context, query *AssetQuery, nodes []*Link, init func(*Link), assign func(*Link, *Asset)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[xid.ID]*Link)
	nids := make(map[xid.ID]map[*Link]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(link.AssetsTable)
		s.Join(joinT).On(s.C(asset.FieldID), joinT.C(link.AssetsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(link.AssetsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(link.AssetsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(xid.ID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*xid.ID)
				inValue := *values[1].(*xid.ID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Link]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Asset](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "assets" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (lq *LinkQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	if len(lq.modifiers) > 0 {
		_spec.Modifiers = lq.modifiers
	}
	_spec.Node.Columns = lq.ctx.Fields
	if len(lq.ctx.Fields) > 0 {
		_spec.Unique = lq.ctx.Unique != nil && *lq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *LinkQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(link.Table, link.Columns, sqlgraph.NewFieldSpec(link.FieldID, field.TypeString))
	_spec.From = lq.sql
	if unique := lq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lq.path != nil {
		_spec.Unique = true
	}
	if fields := lq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, link.FieldID)
		for i := range fields {
			if fields[i] != link.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *LinkQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(link.Table)
	columns := lq.ctx.Fields
	if len(columns) == 0 {
		columns = link.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lq.ctx.Unique != nil && *lq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range lq.modifiers {
		m(selector)
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (lq *LinkQuery) Modify(modifiers ...func(s *sql.Selector)) *LinkSelect {
	lq.modifiers = append(lq.modifiers, modifiers...)
	return lq.Select()
}

// LinkGroupBy is the group-by builder for Link entities.
type LinkGroupBy struct {
	selector
	build *LinkQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LinkGroupBy) Aggregate(fns ...AggregateFunc) *LinkGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the selector query and scans the result into the given value.
func (lgb *LinkGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lgb.build.ctx, "GroupBy")
	if err := lgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LinkQuery, *LinkGroupBy](ctx, lgb.build, lgb, lgb.build.inters, v)
}

func (lgb *LinkGroupBy) sqlScan(ctx context.Context, root *LinkQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lgb.flds)+len(lgb.fns))
		for _, f := range *lgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LinkSelect is the builder for selecting fields of Link entities.
type LinkSelect struct {
	*LinkQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ls *LinkSelect) Aggregate(fns ...AggregateFunc) *LinkSelect {
	ls.fns = append(ls.fns, fns...)
	return ls
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LinkSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ls.ctx, "Select")
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LinkQuery, *LinkSelect](ctx, ls.LinkQuery, ls, ls.inters, v)
}

func (ls *LinkSelect) sqlScan(ctx context.Context, root *LinkQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ls.fns))
	for _, fn := range ls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ls *LinkSelect) Modify(modifiers ...func(s *sql.Selector)) *LinkSelect {
	ls.modifiers = append(ls.modifiers, modifiers...)
	return ls
}
