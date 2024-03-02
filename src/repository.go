package src

import "context"

// URLRepository ...
type URLRepository interface {
	Store(context.Context, URL) Either[UUID]
	Find(context.Context, UUID) Either[URL]
}
