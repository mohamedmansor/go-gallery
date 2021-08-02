package controller

import (
	"blog/api/service"
	"blog/models"
	"blog/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//ImageController -> ImageController
type ImageController struct {
	service service.ImageService
}

//NewImageController : NewImageController
func NewImageController(s service.ImageService) ImageController {
	return ImageController{
		service: s,
	}
}

// GetImages : GetImages controller
func (p ImageController) GetImages(ctx *gin.Context) {
	var images models.Image

	keyword := ctx.Query("keyword")

	data, total, err := p.service.FindAll(images, keyword)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
		return
	}
	respArr := make([]map[string]interface{}, 0, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Image result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

// AddImage : AddImage controller
func (p *ImageController) AddImage(ctx *gin.Context) {
	var image models.Image
	ctx.ShouldBindJSON(&image)

	if image.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if image.ImageUrl == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "ImageUrl is required")
		return
	}
	err := p.service.Save(image)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create image")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Image")
}

//GetImage : get image by id
func (p *ImageController) GetImage(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	var image models.Image
	image.ID = id
	foundImage, err := p.service.Find(image)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Error Finding Image")
		return
	}
	response := foundImage.ResponseMap()

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Image",
		Data:    &response})

}

//DeleteImage : Deletes Image
func (p *ImageController) DeleteImage(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.Delete(id)

	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to delete Image")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	c.JSON(http.StatusOK, response)
}

//UpdateImage : get update by id
func (p ImageController) UpdateImage(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var image models.Image
	image.ID = id

	imageRecord, err := p.service.Find(image)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Image with given id not found")
		return
	}
	ctx.ShouldBindJSON(&imageRecord)

	if imageRecord.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if imageRecord.ImageUrl == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "ImageUrl is required")
		return
	}

	if err := p.service.Update(imageRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Image")
		return
	}
	response := imageRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Image",
		Data:    response,
	})
}
