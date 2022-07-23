package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/cip8/autoname"
)

var (
	Client  *storage.Client
	Handler *storage.BucketHandle
	CTX     context.Context
)

const bucketName = "nicos-first-bucket"

func InitGoogle() {

	var err error
	ctx := context.Background()
	Client, err = storage.NewClient(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer Client.Close()

	Handler = Client.Bucket(bucketName)

}

func GetSignedURL(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")

	opts := &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		Method:         "PUT",
		GoogleAccessID: os.Getenv("CP_EMAIL"),
		PrivateKey:     []byte(os.Getenv("CP_PRIVATE_KEY")),

		Expires: time.Now().Add(15 * time.Minute),
	}

	name := autoname.Generate("image")
	u, err := Handler.SignedURL(name, opts)
	if err != nil {
		fmt.Printf("Bucket(%q).SignedURL: %v", bucketName, err)
	}

	type response struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	}

	res := response{
		URL:  u,
		Name: name,
	}

	json.NewEncoder(w).Encode(res)
}
