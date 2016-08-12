package db

import (
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	defaultLimit = "25"
	defaultPage  = "1"
	defaultOrder = "desc"
)

type Pagination struct {
	Limit  int
	Page   int
	LastID int
	Order  string
}

func (p *Pagination) Paginate(c *gin.Context) (*gorm.DB, error) {
	db := DBInstance(c)
	limitQuery := c.DefaultQuery("limit", defaultLimit)
	pageQuery := c.DefaultQuery("page", defaultPage)
	lastIDQuery := c.Query("last_id")

	p.Order = c.DefaultQuery("order", defaultOrder)

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		return db, err
	}

	p.Limit = int(math.Max(1, math.Min(10000, float64(limit))))

	if lastIDQuery != "" {
		lastID, err := strconv.Atoi(lastIDQuery)
		if err != nil {
			return db, err
		}

		p.LastID = int(math.Max(0, float64(lastID)))

		if p.Order == "asc" {
			return db.Where("id > ?", p.LastID).Limit(p.Limit).Order("id asc"), nil
		}

		return db.Where("id < ?", p.LastID).Limit(p.Limit).Order("id desc"), nil
	}

	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		return db, err
	}

	p.Page = int(math.Max(1, float64(page)))

	return db.Offset(limit * (p.Page - 1)).Limit(p.Limit), nil
}

func (p *Pagination) SetHeaderLink(c *gin.Context, index int) {
	var link string

	if p.LastID != 0 {
		link = fmt.Sprintf("<http://%v%v?limit=%v&last_id=%v&order=%v>; rel=\"next\"", c.Request.Host, c.Request.URL.Path, p.Limit, index, p.Order)
	} else {
		if p.Page == 1 {
			link = fmt.Sprintf("<http://%v%v?limit=%v&page=%v>; rel=\"next\"", c.Request.Host, c.Request.URL.Path, p.Limit, p.Page+1)
		} else {
			link = fmt.Sprintf(
				"<http://%v%v?limit=%v&page=%v>; rel=\"next\",<http://%v%v?limit=%v&page=%v>; rel=\"prev\"",
				c.Request.Host, c.Request.URL.Path, p.Limit, p.Page+1, c.Request.Host, c.Request.URL.Path, p.Limit, p.Page-1,
			)
		}
	}

	c.Header("Link", link)
}
