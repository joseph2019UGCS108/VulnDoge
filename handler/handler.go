package handler

import (
	"net/http"

	"github.com/burpOverflow/VulnDoge/api"
	"github.com/burpOverflow/VulnDoge/csrf"
	"github.com/burpOverflow/VulnDoge/directoryTrversal"
	"github.com/burpOverflow/VulnDoge/loadimage"
)

func HandleDirectoryTrversal() {
	http.HandleFunc("/directoryTrversal/", directoryTrversal.DirectoryTrversalHandler)
	http.HandleFunc("/directoryTrversal/easy/", directoryTrversal.Easy)
	http.HandleFunc("/directoryTrversal/easy2/", directoryTrversal.Easy2)
	http.HandleFunc("/directoryTrversal/medium1/", directoryTrversal.Medium1)
	http.HandleFunc("/directoryTrversal/medium2/", directoryTrversal.Medium2)
	http.HandleFunc("/directoryTrversal/medium3/", directoryTrversal.Medium3)
	http.HandleFunc("/directoryTrversal/medium4/", directoryTrversal.Medium4)
}

func HandleCSRF() {
	http.HandleFunc("/csrf/", csrf.CSRFHandler)

	http.HandleFunc("/csrf/easy1/", csrf.Easy1)
	http.HandleFunc("/csrf/easy1/login/", csrf.LoginEasy1)
	http.HandleFunc("/csrf/easy1/create/", csrf.CreateEasy1)
	http.HandleFunc("/csrf/easy1/myaccount/", csrf.MyAccount)
	http.HandleFunc("/csrf/easy1/logout/", csrf.LogoutEasy1)
	http.HandleFunc("/csrf/easy1/changepassword/", csrf.ChangePassword)

	http.HandleFunc("/csrf/easy2/", csrf.Easy2)
	http.HandleFunc("/csrf/easy2/login/", csrf.LoginEasy2)
	http.HandleFunc("/csrf/easy2/create/", csrf.CreateEasy2)
	http.HandleFunc("/csrf/easy2/myaccount/", csrf.MyAccountEasy2)
	http.HandleFunc("/csrf/easy2/logout/", csrf.LogoutEasy2)
	http.HandleFunc("/csrf/easy2/changepassword/", csrf.ChangePasswordEasy2)

}

func HandleAPI() {
	http.HandleFunc("/api/solution", api.APISolution)
}

func LoadImage() {
	http.HandleFunc("/loadimage", loadimage.LoadImageHandler)
	http.HandleFunc("/img", loadimage.Defult2)
}
