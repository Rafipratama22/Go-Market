package helper

import (
	// "github.com/jinzhu/gorm"
	"gorm.io/gorm"
	"strconv"
	// "github.com/gin-gonic/gin"
)


func Paginate(pagination map[string]string) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
	  page, _ := strconv.Atoi(pagination["page"])
	  if page == 0 {
		page = 1
	  }
  
	  pageSize, _ := strconv.Atoi(pagination["limit"])
	  switch {
	  case pageSize > 100:
		pageSize = 100
	  case pageSize <= 0:
		pageSize = 10
	  }
  
	  offset := (page - 1) * pageSize
	  return db.Offset(offset).Limit(pageSize)
	}
  }