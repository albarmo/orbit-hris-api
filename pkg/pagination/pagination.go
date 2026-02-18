package pagination

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Filter struct {
    Page  int
    Limit int
    Sort  string
}

// Bind reads pagination params from the request query string.
func (f *Filter) Bind(ctx *gin.Context) {
    // defaults
    f.Page = 1
    f.Limit = 10

    if p := ctx.Query("page"); p != "" {
        if v, err := strconv.Atoi(p); err == nil && v > 0 {
            f.Page = v
        }
    }

    if l := ctx.Query("limit"); l != "" {
        if v, err := strconv.Atoi(l); err == nil && v > 0 {
            f.Limit = v
        }
    }

    f.Sort = ctx.Query("sort")
}

type Paginator struct {
    DB    *gorm.DB
    Page  int
    Limit int
    Total int64
}

func NewPaginator(db *gorm.DB, filter *Filter) (*Paginator, error) {
    if filter == nil {
        filter = &Filter{Page: 1, Limit: 10}
    }

    var total int64
    if err := db.Count(&total).Error; err != nil {
        return nil, err
    }

    return &Paginator{
        DB:    db,
        Page:  filter.Page,
        Limit: filter.Limit,
        Total: total,
    }, nil
}

func (p *Paginator) Find(dest interface{}) *gorm.DB {
    if p.Page < 1 {
        p.Page = 1
    }
    if p.Limit <= 0 {
        p.Limit = 10
    }
    offset := (p.Page - 1) * p.Limit
    return p.DB.Limit(p.Limit).Offset(offset).Find(dest)
}

type Page[T any] struct {
    Page       int   `json:"page"`
    Limit      int   `json:"limit"`
    Total      int64 `json:"total"`
    TotalPages int   `json:"total_pages"`
    Data       []T   `json:"data"`
}

func (p *Page[T]) Set(data []T, page, limit int, total int64) {
    p.Page = page
    p.Limit = limit
    p.Total = total
    if limit > 0 {
        p.TotalPages = int(math.Ceil(float64(total) / float64(limit)))
    } else {
        p.TotalPages = 0
    }
    p.Data = data
}
