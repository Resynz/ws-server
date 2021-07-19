/**
 * @Author: Resynz
 * @Date: 2021/7/19 15:12
 */
package tools

import (
	"encoding/json"
	"fmt"
	"github.com/rosbit/go-wget"
	"net/http"
	"ws-server/config"
)

func CheckAuth(rawQuery string) (string, config.PlatformType, error) {
	if config.Conf.AuthUrl == "" {
		return rawQuery, 0, nil
	}
	reqUrl := fmt.Sprintf("%s?%s", config.Conf.AuthUrl, rawQuery)
	method := "GET"
	status, content, _, err := wget.Wget(reqUrl, method, nil, nil)
	if err != nil {
		return "", 0, err
	}

	if status != http.StatusOK {
		return "", 0, fmt.Errorf("check auth bad status:%d", status)
	}

	type res struct {
		Platform config.PlatformType `json:"platform"`
		UserId   string              `json:"user_id"`
	}

	var r res
	if err = json.Unmarshal(content, &r); err != nil {
		return "", 0, err
	}
	return r.UserId, r.Platform, nil
}
