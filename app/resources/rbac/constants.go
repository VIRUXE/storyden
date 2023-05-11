package rbac

type Action = string

const (
	ActionCreate Action = "create"
	ActionRead   Action = "read"
	ActionUpdate Action = "update"
	ActionDelete Action = "delete"
)

type Resource = string

const (
	ResourceAccount Resource = "account"
	ResourceThread  Resource = "thread"
	ResourcePost    Resource = "post"
)