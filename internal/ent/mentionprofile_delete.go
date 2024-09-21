// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/mentionprofile"
	"github.com/Southclaws/storyden/internal/ent/predicate"
)

// MentionProfileDelete is the builder for deleting a MentionProfile entity.
type MentionProfileDelete struct {
	config
	hooks    []Hook
	mutation *MentionProfileMutation
}

// Where appends a list predicates to the MentionProfileDelete builder.
func (mpd *MentionProfileDelete) Where(ps ...predicate.MentionProfile) *MentionProfileDelete {
	mpd.mutation.Where(ps...)
	return mpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mpd *MentionProfileDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, mpd.sqlExec, mpd.mutation, mpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mpd *MentionProfileDelete) ExecX(ctx context.Context) int {
	n, err := mpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mpd *MentionProfileDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(mentionprofile.Table, sqlgraph.NewFieldSpec(mentionprofile.FieldID, field.TypeString))
	if ps := mpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mpd.mutation.done = true
	return affected, err
}

// MentionProfileDeleteOne is the builder for deleting a single MentionProfile entity.
type MentionProfileDeleteOne struct {
	mpd *MentionProfileDelete
}

// Where appends a list predicates to the MentionProfileDelete builder.
func (mpdo *MentionProfileDeleteOne) Where(ps ...predicate.MentionProfile) *MentionProfileDeleteOne {
	mpdo.mpd.mutation.Where(ps...)
	return mpdo
}

// Exec executes the deletion query.
func (mpdo *MentionProfileDeleteOne) Exec(ctx context.Context) error {
	n, err := mpdo.mpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{mentionprofile.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mpdo *MentionProfileDeleteOne) ExecX(ctx context.Context) {
	if err := mpdo.Exec(ctx); err != nil {
		panic(err)
	}
}