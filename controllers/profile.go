package controllers

import (
	"net/http"
	"strings"

	dbpkg "github.com/dtan4/apig-sample/db"
	"github.com/dtan4/apig-sample/helper"
	"github.com/dtan4/apig-sample/models"
	"github.com/dtan4/apig-sample/version"

	"github.com/gin-gonic/gin"
)

func GetProfiles(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	preloads := c.DefaultQuery("preloads", "")
	fields := helper.ParseFields(c.DefaultQuery("fields", "*"))
	ids := c.DefaultQuery("ids", "")

	pagination := dbpkg.Pagination{}
	db, err := pagination.Paginate(c)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db = dbpkg.SetPreloads(preloads, db)

	if ids != "" {
		db = db.Where("id IN (?)", strings.Split(ids, ","))
	}

	var profiles []models.Profile
	if err := db.Select("*").Find(&profiles).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// paging
	var index int
	if len(profiles) < 1 {
		index = 0
	} else {
		index = int(profiles[len(profiles)-1].ID)
	}
	pagination.SetHeaderLink(c, index)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	fieldMap := []map[string]interface{}{}
	for _, profile := range profiles {
		fieldMap = append(fieldMap, helper.FieldToMap(profile, fields))
	}
	c.JSON(200, fieldMap)
}

func GetProfile(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Params.ByName("id")
	preloads := c.DefaultQuery("preloads", "")
	fields := helper.ParseFields(c.DefaultQuery("fields", "*"))

	db := dbpkg.DBInstance(c)
	db = dbpkg.SetPreloads(preloads, db)

	var profile models.Profile
	if err := db.Select("*").First(&profile, id).Error; err != nil {
		content := gin.H{"error": "profile with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	fieldMap := helper.FieldToMap(profile, fields)
	c.JSON(200, fieldMap)
}

func CreateProfile(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	var profile models.Profile
	c.Bind(&profile)
	if db.Create(&profile).Error != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(201, profile)
}

func UpdateProfile(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	var profile models.Profile
	if db.First(&profile, id).Error != nil {
		content := gin.H{"error": "profile with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.Bind(&profile)
	db.Save(&profile)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(200, profile)
}

func DeleteProfile(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	var profile models.Profile
	if db.First(&profile, id).Error != nil {
		content := gin.H{"error": "profile with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	db.Delete(&profile)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
