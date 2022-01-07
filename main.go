package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mrrizal/devcode-backend-challenge-fasthttp/configs"
	"github.com/mrrizal/devcode-backend-challenge-fasthttp/database"
	"github.com/mrrizal/devcode-backend-challenge-fasthttp/handler"
	"github.com/mrrizal/devcode-backend-challenge-fasthttp/utils"
	"github.com/valyala/fasthttp"
)

var Config configs.Conf

func loadEnv() {
	Config.MysqlHost = os.Getenv("MYSQL_HOST")
	Config.MysqlUser = os.Getenv("MYSQL_USER")
	Config.MysqlPassword = os.Getenv("MYSQL_PASSWORD")
	Config.MysqlDBName = os.Getenv("MYSQL_DBNAME")
	Config.Log = false
	log, err := strconv.ParseBool(os.Getenv("LOG"))
	if err == nil {
		Config.Log = log
	}
	Config.Port = 3030
}

func routes(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	if strings.Contains(path, "activity-groups") {
		activityGroupsMap := make(map[string]func(ctx *fasthttp.RequestCtx))
		activityGroupsMap["GET"] = handler.GetActivityGroups
		activityGroupsMap["POST"] = handler.CreateActivityGroups
		activityGroupsMap["PATCH"] = handler.PatchActivityGroups
		activityGroupsMap["DELETE"] = handler.DeleteActivityGroups

		if handler, ok := activityGroupsMap[string(ctx.Method())]; ok {
			handler(ctx)
		} else {
			utils.ResponseHandler(ctx, fasthttp.StatusForbidden,
				utils.GenerateErrorMessage("Forbidden", "Forbidden", nil))
		}

	} else if strings.Contains(path, "todo-items") {
		fmt.Println("haha")
	}
}

func main() {
	loadEnv()

	listenAddr := fmt.Sprintf("0.0.0.0:%d", Config.Port)

	if err := database.InitDatabase(Config); err != nil {
		log.Fatal(err.Error())
	}

	requestHandler := routes

	if err := fasthttp.ListenAndServe(listenAddr, requestHandler); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
