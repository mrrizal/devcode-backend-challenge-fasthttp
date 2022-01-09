package handler

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mrrizal/devcode-backend-challenge-fasthttp/models"
	"github.com/mrrizal/devcode-backend-challenge-fasthttp/utils"
	"github.com/valyala/fasthttp"
)

func CreateActivityGroups(ctx *fasthttp.RequestCtx) {
	var activityGroup models.ActivityGroup

	if err := json.Unmarshal(ctx.PostBody(), &activityGroup); err != nil {
		utils.ResponseHandler(ctx, fasthttp.StatusBadRequest,
			utils.GenerateResponse("Bad Request", err.Error(), nil))
		return
	}

	if err := activityGroup.Validate(); err != nil {
		utils.ResponseHandler(ctx, fasthttp.StatusBadRequest,
			utils.GenerateResponse("Bad Request", err.Error(), nil))
		return
	}

	now := time.Now()
	activityGroup.CreatedAt = now
	activityGroup.UpdatedAt = now
	insertedID, err := activityGroup.Insert()

	if err != nil {
		utils.ResponseHandler(ctx, fasthttp.StatusInternalServerError,
			utils.GenerateResponse("Internal Server Error", err.Error(), nil))
		return
	}

	activityGroup.ID = insertedID
	utils.ResponseHandler(ctx, fasthttp.StatusCreated,
		utils.GenerateResponse("Success", "Success", activityGroup))
}

func parseActivityGroupID(ctx *fasthttp.RequestCtx, splitedPath []string) (int, error) {
	activityGroupID, err := strconv.Atoi(splitedPath[len(splitedPath)-1])
	if err != nil {
		utils.ResponseHandler(ctx, fasthttp.StatusNotFound,
			utils.GenerateResponse("Not Found", fmt.Sprintf("Activity with ID %s Not Found",
				splitedPath[len(splitedPath)-1]), nil))
		return 0, err
	}
	return activityGroupID, err
}

func GetActivityGroupsByID(ctx *fasthttp.RequestCtx, splitedPath []string) {
	activityGroupID, err := parseActivityGroupID(ctx, splitedPath)
	if err != nil {
		return
	}
	activityGroup := new(models.ActivityGroup)
	if err := activityGroup.GetByID(activityGroupID, activityGroup); err != nil {
		utils.ResponseHandler(ctx, fasthttp.StatusNotFound,
			utils.GenerateResponse("Not Found", fmt.Sprintf("Activity with ID %d Not Found",
				activityGroupID), nil))
		return
	}
	utils.ResponseHandler(ctx, fasthttp.StatusOK,
		utils.GenerateResponse("Success", "Success", activityGroup))
}

func GetAllActivityGroups(ctx *fasthttp.RequestCtx) {
	activityGroups := []models.ActivityGroup{}

	activityGroup := new(models.ActivityGroup)
	if err := activityGroup.GetAll(&activityGroups); err != nil {
		utils.ResponseHandler(ctx, fasthttp.StatusInternalServerError,
			utils.GenerateResponse("Internal Server Error", err.Error(), nil))
		return
	}
	utils.ResponseHandler(ctx, fasthttp.StatusOK,
		utils.GenerateResponse("Success", "Success", activityGroups))
}

func GetActivityGroups(ctx *fasthttp.RequestCtx) {
	path := strings.TrimSuffix(string(ctx.Path()), "/")
	splitedPath := strings.Split(path, "/")

	// get activity group by id
	if len(splitedPath) == 3 {
		GetActivityGroupsByID(ctx, splitedPath)
		return
	}

	// get all activity group
	GetAllActivityGroups(ctx)
}

func PatchActivityGroups(ctx *fasthttp.RequestCtx) {
	path := strings.TrimSuffix(string(ctx.Path()), "/")
	splitedPath := strings.Split(path, "/")

	if len(splitedPath) == 3 {
		var tempActivityGroup models.ActivityGroup

		if err := json.Unmarshal(ctx.PostBody(), &tempActivityGroup); err != nil {
			utils.ResponseHandler(ctx, fasthttp.StatusBadRequest,
				utils.GenerateResponse("Bad Request", err.Error(), nil))
			return
		}

		if err := tempActivityGroup.Validate(); err != nil {
			utils.ResponseHandler(ctx, fasthttp.StatusBadRequest,
				utils.GenerateResponse("Bad Request", err.Error(), nil))
			return
		}

		activityGroupID, err := parseActivityGroupID(ctx, splitedPath)
		if err != nil {
			return
		}
		tempActivityGroup.ID = activityGroupID

		rowsAffected, err := tempActivityGroup.Update()
		if err != nil {
			utils.ResponseHandler(ctx, fasthttp.StatusInternalServerError,
				utils.GenerateResponse("Internal Server Error", err.Error(), nil))
			return
		}

		if rowsAffected == 0 {
			utils.ResponseHandler(ctx, fasthttp.StatusNotFound,
				utils.GenerateResponse("Not Found", fmt.Sprintf("Activity with ID %s Not Found",
					splitedPath[len(splitedPath)-1]), nil))
			return
		}

		var activityGroup models.ActivityGroup
		activityGroup.GetByID(activityGroupID, &activityGroup)

		utils.ResponseHandler(ctx, fasthttp.StatusOK,
			utils.GenerateResponse("Success", "Success", activityGroup))
		return
	}
	utils.ResponseHandler(ctx, fasthttp.StatusForbidden,
		utils.GenerateResponse("Forbidden", "Forbidden", nil))
}

func DeleteActivityGroups(ctx *fasthttp.RequestCtx) {
	path := strings.TrimSuffix(string(ctx.Path()), "/")
	splitedPath := strings.Split(path, "/")

	// get activity group by id
	if len(splitedPath) == 3 {
		activityGroupID, err := parseActivityGroupID(ctx, splitedPath)
		if err != nil {
			return
		}

		var activityGroup models.ActivityGroup
		activityGroup.ID = activityGroupID

		rowsAffected, err := activityGroup.Delete()
		if err != nil {
			utils.ResponseHandler(ctx, fasthttp.StatusInternalServerError,
				utils.GenerateResponse("Internal Server Error", err.Error(), nil))
			return
		}

		if rowsAffected == 0 {
			utils.ResponseHandler(ctx, fasthttp.StatusNotFound,
				utils.GenerateResponse("Not Found", fmt.Sprintf("Activity with ID %s Not Found",
					splitedPath[len(splitedPath)-1]), nil))
			return
		}

		utils.ResponseHandler(ctx, fasthttp.StatusOK,
			utils.GenerateResponse("Success", "Success", nil))
		return
	}
	utils.ResponseHandler(ctx, fasthttp.StatusForbidden,
		utils.GenerateResponse("Forbidden", "Forbidden", nil))
}
