// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package docs

import (
	sheets "google.golang.org/api/sheets/v4"
	"net/http"
	"log"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"os"
	"net/url"
	"encoding/json"
	"fmt"
)
const spreadsheetId = "1AZaTtLAlKZuh_zBASLj2rHA2ux0jqOVxBVwwRYMi7CE"
const ClientID 	    = "165049723351-mgcbnem17vt14plfhtbfdcerc1ona2p7.apps.googleusercontent.com"
const authCode  =  "4/H7iL6R6BSstU5-W0V7WgI9cPZttAjOzHH5pEmwYS8UQ#"

type SheetsGoogleDocs struct {
	Service *sheets.Service
}
// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("sheets.googleapis.com-go-quickstart.json")), err
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	//authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	//fmt.Printf("Go to the following link in your browser then type the "+
	//	"authorization code: \n%v\n", authURL)

	//var code string
	//if _, err := fmt.Scan(&code); err != nil {
	//	log.Fatalf("Unable to read authorization code %v", err)
	//}

	tok, err := config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}
// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}
// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}
	return config.Client(ctx, tok)
}
func newClient() *http.Client{
	ctx := context.Background()
	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/sheets.googleapis.com-go-quickstart.json
	b, err := ioutil.ReadFile("config/oauth2.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	return getClient(ctx, config)
}
func (sheet * SheetsGoogleDocs) Init() (err error){
	if sheet.Service, err = sheets.New(newClient()); err != nil {
		return err
	}

	return nil
}

func (sheet * SheetsGoogleDocs) Read() ( *sheets.ValueRange, error) {
	sh := sheet.Service.Spreadsheets.Values

	readRange := "A1:A1000"
	doGet := sh.Get(spreadsheetId, readRange)
	resp, err := doGet.Do();
	if err != nil {
		return nil, err
	}
	return resp, nil
}
