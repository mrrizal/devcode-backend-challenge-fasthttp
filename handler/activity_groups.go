package handler

import (
	"encoding/json"
	"time"

	"github.com/mrrizal/devcode-backend-challenge-fasthttp/database"
	"github.com/mrrizal/devcode-backend-challenge-fasthttp/models"
	"github.com/mrrizal/devcode-backend-challenge-fasthttp/utils"
	"github.com/valyala/fasthttp"
)

func CreateActivityGroups(ctx *fasthttp.RequestCtx) {
	var activityGroup models.ActivityGroup
	db := database.DBConn

	if err := json.Unmarshal(ctx.PostBody(), &activityGroup); err != nil {
		utils.ResponseHandler(ctx, fasthttp.StatusBadRequest,
			utils.GenerateErrorMessage("Bad Request", err.Error(), nil))
		return
	}

	if err := activityGroup.Validate(); err != nil {
		utils.ResponseHandler(ctx, fasthttp.StatusBadRequest,
			utils.GenerateErrorMessage("Bad Request", err.Error(), nil))
		return
	}

	stmt, err := db.Prepare("insert into activities (email, title, created_at, updated_at) values (?, ?, ?, ?)")
	if err != nil {
		utils.ResponseHandler(ctx, fasthttp.StatusBadRequest, map[string]interface{}{"message": err.Error()})
		return
	}
	defer stmt.Close()

	now := time.Now()
	activityGroup.CreatedAt = now
	activityGroup.UpdatedAt = now
	resp, err := stmt.Exec(activityGroup.Email, activityGroup.Title, activityGroup.CreatedAt, activityGroup.UpdatedAt)
	if err != nil {
		utils.ResponseHandler(ctx, fasthttp.StatusInternalServerError,
			utils.GenerateErrorMessage("Internal Server Error", err.Error(), nil))
		return
	}

	insertedID, err := resp.LastInsertId()
	if err != nil {
		utils.ResponseHandler(ctx, fasthttp.StatusInternalServerError,
			utils.GenerateErrorMessage("Internal Server Error", err.Error(), nil))
		return
	}

	activityGroup.ID = int(insertedID)
	utils.ResponseHandler(ctx, fasthttp.StatusCreated,
		utils.GenerateErrorMessage("Success", "Success", activityGroup))
}

func GetActivityGroups(ctx *fasthttp.RequestCtx) {

}

func PatchActivityGroups(ctx *fasthttp.RequestCtx) {
}

func DeleteActivityGroups(ctx *fasthttp.RequestCtx) {
}
