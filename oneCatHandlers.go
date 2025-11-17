package main

import "net/http"

func getCat(req *http.Request) (int, any) {
	catId := req.PathValue("catId")
	Logger.Info("Getting cat", catId)

	if cat , found := catsDatabase[catId]; found {
		Logger.Info("Cat found")
		return http.StatusOK, cat
	} else {
		Logger.Info("Cat not found")
		return http.StatusNotFound, "Cat not found"
	}
}

func deleteCat(req *http.Request) (int, any) {
	catId := req.PathValue("catId")
	Logger.Info("Deleting cat", catId)

	if _, found := catsDatabase[catId]; found {
		delete(catsDatabase, catId)
		Logger.Info("Cat deleted")
		return http.StatusNoContent, catId
	} else {
		Logger.Info("Cat not found")
		return http.StatusNotFound, "Cat not found"
	}
}