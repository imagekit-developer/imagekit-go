// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/imagekit-developer/imagekit-go/v2/internal/apijson"
	"github.com/imagekit-developer/imagekit-go/v2/internal/requestconfig"
	"github.com/imagekit-developer/imagekit-go/v2/option"
	"github.com/imagekit-developer/imagekit-go/v2/packages/param"
	"github.com/imagekit-developer/imagekit-go/v2/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type TagCursor[T any] struct {
	Tags       []T    `json:"tags"`
	NextCursor string `json:"nextCursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tags        respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r TagCursor[T]) RawJSON() string { return r.JSON.raw }
func (r *TagCursor[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *TagCursor[T]) GetNextPage() (res *TagCursor[T], err error) {
	if len(r.Tags) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *TagCursor[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &TagCursor[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type TagCursorAutoPager[T any] struct {
	page *TagCursor[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewTagCursorAutoPager[T any](page *TagCursor[T], err error) *TagCursorAutoPager[T] {
	return &TagCursorAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *TagCursorAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Tags) == 0 {
		return false
	}
	if r.idx >= len(r.page.Tags) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Tags) == 0 {
			return false
		}
	}
	r.cur = r.page.Tags[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *TagCursorAutoPager[T]) Current() T {
	return r.cur
}

func (r *TagCursorAutoPager[T]) Err() error {
	return r.err
}

func (r *TagCursorAutoPager[T]) Index() int {
	return r.run
}
