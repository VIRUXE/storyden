package thread

import (
	"time"

	"github.com/Southclaws/dt"
	"github.com/Southclaws/fault"
	"github.com/Southclaws/opt"

	"github.com/Southclaws/storyden/app/resources/asset"
	"github.com/Southclaws/storyden/app/resources/collection"
	"github.com/Southclaws/storyden/app/resources/datagraph"
	"github.com/Southclaws/storyden/app/resources/link/link_ref"
	"github.com/Southclaws/storyden/app/resources/post"
	"github.com/Southclaws/storyden/app/resources/post/category"
	"github.com/Southclaws/storyden/app/resources/post/reaction"
	"github.com/Southclaws/storyden/app/resources/post/reply"
	"github.com/Southclaws/storyden/app/resources/profile"
	"github.com/Southclaws/storyden/app/resources/tag/tag_ref"
	"github.com/Southclaws/storyden/app/resources/visibility"
	"github.com/Southclaws/storyden/internal/ent"
)

type Thread struct {
	post.Post

	Title  string
	Slug   string
	Short  string
	Pinned bool

	ReplyStatus post.ReplyStatus
	Replies     []*reply.Reply
	Category    category.Category
	Visibility  visibility.Visibility
	Tags        tag_ref.Tags
	Collections []*collection.Collection
	Related     datagraph.ItemList
}

func (*Thread) GetResourceName() string { return "thread" }

func (t *Thread) GetName() string       { return t.Title }
func (t *Thread) GetSlug() string       { return t.Slug }
func (t *Thread) GetDesc() string       { return t.Short }
func (t *Thread) GetCreated() time.Time { return t.CreatedAt }
func (t *Thread) GetUpdated() time.Time { return t.UpdatedAt }

func FromModel(ls post.PostLikesMap, rs post.PostRepliesMap) func(m *ent.Post) (*Thread, error) {
	return func(m *ent.Post) (*Thread, error) {
		categoryEdge, err := m.Edges.CategoryOrErr()
		if err != nil {
			return nil, fault.Wrap(err)
		}

		authorEdge, err := m.Edges.AuthorOrErr()
		if err != nil {
			return nil, fault.Wrap(err)
		}

		category := category.FromModel(categoryEdge)

		pro, err := profile.ProfileFromModel(authorEdge)
		if err != nil {
			return nil, fault.Wrap(err)
		}

		replies, err := dt.MapErr(m.Edges.Posts, reply.FromModel(ls))
		if err != nil {
			return nil, fault.Wrap(err)
		}

		collectionsEdge := opt.NewIf(m.Edges.Collections, func(c []*ent.Collection) bool { return c != nil })

		collections, err := opt.MapErr(collectionsEdge, func(c []*ent.Collection) ([]*collection.Collection, error) {
			out, err := dt.MapErr(c, collection.MapCollection)
			if err != nil {
				return nil, fault.Wrap(err)
			}
			return out, nil
		})
		if err != nil {
			return nil, fault.Wrap(err)
		}

		link := opt.Map(opt.NewPtr(m.Edges.Link), func(in ent.Link) link_ref.LinkRef {
			return *link_ref.Map(&in)
		})

		content, err := datagraph.NewRichText(m.Body)
		if err != nil {
			return nil, fault.Wrap(err)
		}

		reacts, err := reaction.MapList(m.Edges.Reacts)
		if err != nil {
			return nil, err
		}

		tags := dt.Map(m.Edges.Tags, tag_ref.Map(nil))

		return &Thread{
			Post: post.Post{
				ID: post.ID(m.ID),

				Content: content,
				Author:  *pro,
				Likes:   ls.Status(m.ID),
				Reacts:  reacts,
				Assets:  dt.Map(m.Edges.Assets, asset.FromModel),
				WebLink: link,
				Meta:    m.Metadata,

				CreatedAt: m.CreatedAt,
				UpdatedAt: m.UpdatedAt,
				DeletedAt: opt.NewPtr(m.DeletedAt),
			},

			Title:  m.Title,
			Slug:   m.Slug,
			Short:  m.Short,
			Pinned: m.Pinned,

			ReplyStatus: rs.Status(m.ID),
			Replies:     replies,
			Category:    *category,
			Visibility:  visibility.NewVisibilityFromEnt(m.Visibility),
			Tags:        tags,
			Collections: collections.OrZero(),
		}, nil
	}
}
