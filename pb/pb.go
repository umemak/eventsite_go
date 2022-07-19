package pb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const PB_URL = "http://pocketbase:8090"

func CreateUser(email string, password string, passwordConfirm string) ([]byte, error) {
	url := PB_URL + "/api/users"
	jsonString := fmt.Sprintf(`{"email": "%s", "password": "%s", "passwordConfirm": "%s"}`,
		email, password, passwordConfirm)
	// fmt.Printf("%+v\n", jsonString)
	req, err := http.NewRequest("POST", url, strings.NewReader(jsonString))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}
	st := resp.StatusCode
	if st != http.StatusOK {
		return body, fmt.Errorf("http status %d", st)
	}
	return body, nil
}

func AuthViaEmail(email string, password string) ([]byte, error) {
	url := PB_URL + "/api/users/auth-via-email"
	jsonString := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password)
	req, err := http.NewRequest("POST", url, strings.NewReader(jsonString))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}
	st := resp.StatusCode
	if st != http.StatusOK {
		return body, fmt.Errorf("http status %d", st)
	}
	return body, nil
}
