package domain

import (
	"context"
	"encoding/json"
	"net/http"
)

type SnowflakeService struct {
	snowflakeUrl string
}

func NewSnowflakeService(snowflakeUrl string) *SnowflakeService {
	return &SnowflakeService{snowflakeUrl: snowflakeUrl}
}

func (s *SnowflakeService) Generate(ctx context.Context) (int64, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.snowflakeUrl, nil)
	if err != nil {
		return 0, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	jsonRes := &struct{ Id int64 }{}
	err = json.NewDecoder(res.Body).Decode(jsonRes)
	if err != nil {
		return 0, err
	}

	return jsonRes.Id, nil
}
