package handlers

import (
	"encoding/json"
	"fmt"
	"goserver/models"
	"net/http"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"golang.org/x/crypto/bcrypt"
)

var myService *cloudantv1.CloudantV1
var database = "users"

func init() {
	auth := &core.IamAuthenticator{
		ApiKey: "xdF3p0dLcyuooAGRb8ZOQxH1buMP9x5i8lY-bgMFBOA5",
	}

	options := &cloudantv1.CloudantV1Options{
		Authenticator: auth,
		URL:           "https://apikey-v2-30hw28jg73of98ig13usldgtgtcu1qlcsd1mt4xuigx5:89a2a2f53bc81a1e41fad24809547e41@1f487807-3cc5-4ebb-a502-c791c9eb4b78-bluemix.cloudantnosqldb.appdomain.cloud",
	}

	var err error
	myService, err = cloudantv1.NewCloudantV1(options)
	if err != nil {
		panic(err)
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.SignUpData

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user.Username)

	existing, _, _ := myService.GetDocument(myService.NewGetDocumentOptions(database, user.Username))
	if existing != nil {
		var res models.Response
		res.Response = "EX"
		res.Data = ""

		resJson, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resJson)
		return
	}

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	userID := user.Username
	userToAdd := cloudantv1.Document{
		ID: &userID,
	}

	userToAdd.SetProperty("email", user.Email)
	userToAdd.SetProperty("password", string(hashedpassword))

	options := myService.NewPostDocumentOptions(database).SetDocument(&userToAdd)

	creationResponse, _, err := myService.PostDocument(options)
	if err != nil {
		panic(err)
	}

	var res models.Response
	res.Response = "OK"
	res.Data = *creationResponse.Rev

	resJson, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	fmt.Println("GG !")
	w.Header().Set("Content-Type", "application/json")
	w.Write(resJson)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.LoginData

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	retrievedUser, _, err := myService.GetDocument(myService.NewGetDocumentOptions(database, user.Username))
	if err != nil {
		panic(err)
	}

	jsonData, err := retrievedUser.MarshalJSON()
	if err != nil {
		panic(err)
	}

	var verfiy models.Verify

	err = json.Unmarshal(jsonData, &verfiy)
	if err != nil {
		panic(err)
	}

	fmt.Println(verfiy.Id)

	err = bcrypt.CompareHashAndPassword([]byte(verfiy.Password), []byte(user.Password))
	if err != nil {
		var res models.Response
		res.Response = "PI"
		res.Data = ""
		resJson, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resJson)
		return
	} else {
		var res models.Response
		res.Response = "OK"
		res.Data = verfiy.Id

		resJson, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resJson)
		return
	}
}
