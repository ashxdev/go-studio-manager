package handlers

import (
	"go-dls/templates/home"
	"net/http"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	//user := getAuthenticatedUser(r)
	// account, err := db.GetAccountByUserID(user.ID)
	// if err != nil {
	// 	return err
	// }
	//fmt.Printf("%+v\n", user.Account)
	return home.Index().Render(r.Context(), w)
}