package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

//lint:ignore U1000 Ignore unused function temporarily for debugging
type AuthDto struct {
	tableName struct{}  `pg:"user.app_user,alias:au"`
	Id        uuid.UUID `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Name      string    `json:"name"`
}

func (d *DataConn) PingAuthService(token string) (*AuthDto, error) {
	client := &http.Client{}

	url := fmt.Sprintf("%v/ping", d.Env.AuthAddr)

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		return nil, errors.New(string(b))
	}

	responseBody := AuthDto{}

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return nil, err
	}

	return &responseBody, nil
}
