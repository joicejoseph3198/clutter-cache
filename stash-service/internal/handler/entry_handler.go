package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"joicejoseph.dev/clutter-cache/stash-service/internal/data"
	"joicejoseph.dev/clutter-cache/stash-service/internal/service"
)


type EntryHandler struct{
	entryService service.EntryService
}

func NewEntryHandler(entryService service.EntryService) *EntryHandler{
	return &EntryHandler{entryService: entryService}
}

func (h *EntryHandler) FetchAllEntries() gin.HandlerFunc{
	return func(c *gin.Context){
		response := data.WebResponse{
			Status: true,
			Message: "Request processed successfully",
			Data: nil,
		}
	
		entries, err := h.entryService.GetEntries()
		if err!=nil || len(entries) == 0{
			response.Status = false
			response.Message = "No entry found"
			c.JSON(http.StatusNotFound, response)
			return
		}
		response.Data = entries;
		c.JSON(http.StatusOK, response)
	}
	
}