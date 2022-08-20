// Code generated by entc, DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Southclaws/storyden/backend/internal/infrastructure/db/model/account"
	"github.com/Southclaws/storyden/backend/internal/infrastructure/db/model/category"
	"github.com/Southclaws/storyden/backend/internal/infrastructure/db/model/post"
	"github.com/google/uuid"
)

// Post is the model entity for the Post schema.
type Post struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// First holds the value of the "first" field.
	First bool `json:"first,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Pinned holds the value of the "pinned" field.
	Pinned bool `json:"pinned,omitempty"`
	// RootPostID holds the value of the "root_post_id" field.
	RootPostID uuid.UUID `json:"root_post_id,omitempty"`
	// ReplyToPostID holds the value of the "reply_to_post_id" field.
	ReplyToPostID uuid.UUID `json:"reply_to_post_id,omitempty"`
	// Body holds the value of the "body" field.
	Body string `json:"body,omitempty"`
	// Short holds the value of the "short" field.
	Short string `json:"short,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// DeletedAt holds the value of the "deletedAt" field.
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// CategoryID holds the value of the "category_id" field.
	CategoryID uuid.UUID `json:"category_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PostQuery when eager-loading is set.
	Edges         PostEdges `json:"edges"`
	account_posts *uuid.UUID
}

// PostEdges holds the relations/edges for other nodes in the graph.
type PostEdges struct {
	// Author holds the value of the author edge.
	Author *Account `json:"author,omitempty"`
	// Category holds the value of the category edge.
	Category *Category `json:"category,omitempty"`
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// Root holds the value of the root edge.
	Root *Post `json:"root,omitempty"`
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// ReplyTo holds the value of the replyTo edge.
	ReplyTo *Post `json:"replyTo,omitempty"`
	// Replies holds the value of the replies edge.
	Replies []*Post `json:"replies,omitempty"`
	// Reacts holds the value of the reacts edge.
	Reacts []*React `json:"reacts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) AuthorOrErr() (*Account, error) {
	if e.loadedTypes[0] {
		if e.Author == nil {
			// The edge author was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: account.Label}
		}
		return e.Author, nil
	}
	return nil, &NotLoadedError{edge: "author"}
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) CategoryOrErr() (*Category, error) {
	if e.loadedTypes[1] {
		if e.Category == nil {
			// The edge category was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: category.Label}
		}
		return e.Category, nil
	}
	return nil, &NotLoadedError{edge: "category"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[2] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// RootOrErr returns the Root value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) RootOrErr() (*Post, error) {
	if e.loadedTypes[3] {
		if e.Root == nil {
			// The edge root was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: post.Label}
		}
		return e.Root, nil
	}
	return nil, &NotLoadedError{edge: "root"}
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[4] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// ReplyToOrErr returns the ReplyTo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) ReplyToOrErr() (*Post, error) {
	if e.loadedTypes[5] {
		if e.ReplyTo == nil {
			// The edge replyTo was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: post.Label}
		}
		return e.ReplyTo, nil
	}
	return nil, &NotLoadedError{edge: "replyTo"}
}

// RepliesOrErr returns the Replies value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) RepliesOrErr() ([]*Post, error) {
	if e.loadedTypes[6] {
		return e.Replies, nil
	}
	return nil, &NotLoadedError{edge: "replies"}
}

// ReactsOrErr returns the Reacts value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) ReactsOrErr() ([]*React, error) {
	if e.loadedTypes[7] {
		return e.Reacts, nil
	}
	return nil, &NotLoadedError{edge: "reacts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Post) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case post.FieldFirst, post.FieldPinned:
			values[i] = new(sql.NullBool)
		case post.FieldTitle, post.FieldSlug, post.FieldBody, post.FieldShort:
			values[i] = new(sql.NullString)
		case post.FieldCreatedAt, post.FieldUpdatedAt, post.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case post.FieldID, post.FieldRootPostID, post.FieldReplyToPostID, post.FieldCategoryID:
			values[i] = new(uuid.UUID)
		case post.ForeignKeys[0]: // account_posts
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Post", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Post fields.
func (po *Post) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case post.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				po.ID = *value
			}
		case post.FieldFirst:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field first", values[i])
			} else if value.Valid {
				po.First = value.Bool
			}
		case post.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				po.Title = value.String
			}
		case post.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				po.Slug = value.String
			}
		case post.FieldPinned:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field pinned", values[i])
			} else if value.Valid {
				po.Pinned = value.Bool
			}
		case post.FieldRootPostID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field root_post_id", values[i])
			} else if value != nil {
				po.RootPostID = *value
			}
		case post.FieldReplyToPostID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field reply_to_post_id", values[i])
			} else if value != nil {
				po.ReplyToPostID = *value
			}
		case post.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				po.Body = value.String
			}
		case post.FieldShort:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field short", values[i])
			} else if value.Valid {
				po.Short = value.String
			}
		case post.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				po.CreatedAt = value.Time
			}
		case post.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				po.UpdatedAt = value.Time
			}
		case post.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deletedAt", values[i])
			} else if value.Valid {
				po.DeletedAt = new(time.Time)
				*po.DeletedAt = value.Time
			}
		case post.FieldCategoryID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field category_id", values[i])
			} else if value != nil {
				po.CategoryID = *value
			}
		case post.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field account_posts", values[i])
			} else if value.Valid {
				po.account_posts = new(uuid.UUID)
				*po.account_posts = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryAuthor queries the "author" edge of the Post entity.
func (po *Post) QueryAuthor() *AccountQuery {
	return (&PostClient{config: po.config}).QueryAuthor(po)
}

// QueryCategory queries the "category" edge of the Post entity.
func (po *Post) QueryCategory() *CategoryQuery {
	return (&PostClient{config: po.config}).QueryCategory(po)
}

// QueryTags queries the "tags" edge of the Post entity.
func (po *Post) QueryTags() *TagQuery {
	return (&PostClient{config: po.config}).QueryTags(po)
}

// QueryRoot queries the "root" edge of the Post entity.
func (po *Post) QueryRoot() *PostQuery {
	return (&PostClient{config: po.config}).QueryRoot(po)
}

// QueryPosts queries the "posts" edge of the Post entity.
func (po *Post) QueryPosts() *PostQuery {
	return (&PostClient{config: po.config}).QueryPosts(po)
}

// QueryReplyTo queries the "replyTo" edge of the Post entity.
func (po *Post) QueryReplyTo() *PostQuery {
	return (&PostClient{config: po.config}).QueryReplyTo(po)
}

// QueryReplies queries the "replies" edge of the Post entity.
func (po *Post) QueryReplies() *PostQuery {
	return (&PostClient{config: po.config}).QueryReplies(po)
}

// QueryReacts queries the "reacts" edge of the Post entity.
func (po *Post) QueryReacts() *ReactQuery {
	return (&PostClient{config: po.config}).QueryReacts(po)
}

// Update returns a builder for updating this Post.
// Note that you need to call Post.Unwrap() before calling this method if this Post
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Post) Update() *PostUpdateOne {
	return (&PostClient{config: po.config}).UpdateOne(po)
}

// Unwrap unwraps the Post entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Post) Unwrap() *Post {
	tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("model: Post is not a transactional entity")
	}
	po.config.driver = tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Post) String() string {
	var builder strings.Builder
	builder.WriteString("Post(")
	builder.WriteString(fmt.Sprintf("id=%v", po.ID))
	builder.WriteString(", first=")
	builder.WriteString(fmt.Sprintf("%v", po.First))
	builder.WriteString(", title=")
	builder.WriteString(po.Title)
	builder.WriteString(", slug=")
	builder.WriteString(po.Slug)
	builder.WriteString(", pinned=")
	builder.WriteString(fmt.Sprintf("%v", po.Pinned))
	builder.WriteString(", root_post_id=")
	builder.WriteString(fmt.Sprintf("%v", po.RootPostID))
	builder.WriteString(", reply_to_post_id=")
	builder.WriteString(fmt.Sprintf("%v", po.ReplyToPostID))
	builder.WriteString(", body=")
	builder.WriteString(po.Body)
	builder.WriteString(", short=")
	builder.WriteString(po.Short)
	builder.WriteString(", createdAt=")
	builder.WriteString(po.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(po.UpdatedAt.Format(time.ANSIC))
	if v := po.DeletedAt; v != nil {
		builder.WriteString(", deletedAt=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", category_id=")
	builder.WriteString(fmt.Sprintf("%v", po.CategoryID))
	builder.WriteByte(')')
	return builder.String()
}

// Posts is a parsable slice of Post.
type Posts []*Post

func (po Posts) config(cfg config) {
	for _i := range po {
		po[_i].config = cfg
	}
}
