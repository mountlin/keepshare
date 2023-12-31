// Copyright 2023 The KeepShare Authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/KeepShareOrg/keepshare/config"
	lk "github.com/KeepShareOrg/keepshare/pkg/link"
)

// VerifyRecaptchaToken verify google recaptcha token
func VerifyRecaptchaToken(token string) bool {
	secret := config.GoogleRecaptchaSecret()
	verifyApi := "https://www.google.com/recaptcha/api/siteverify"
	verifyUrl := fmt.Sprintf("%s?secret=%s&response=%s", verifyApi, secret, token)
	resp, err := http.Post(verifyUrl, "application/x-www-form-urlencoded", nil)
	if err != nil {
		return false
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	data := map[string]interface{}{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return false
	}
	return data["success"].(bool)
}

func makeKeepSharingLink(channel, originalLink string) string {
	return fmt.Sprintf("https://%s/%s/%s", config.RootDomain(), channel, url.QueryEscape(originalLink))
}

func getOriginalLinks(src []string) (original, invalid []string) {
	original = make([]string, 0, len(src))
	for _, raw := range src {
		raw = strings.TrimSpace(raw)

		// check url
		u, err := url.Parse(raw)
		if err != nil || u.Scheme == "" {
			invalid = append(invalid, raw)
			continue
		}

		// is auto sharing link
		if u.Host == config.RootDomain() {
			_, link, ok := getChannelAndLinkFromURL(u)
			if !ok {
				invalid = append(invalid, raw)
				continue
			}
			raw = link
		}

		original = append(original, lk.Simplify(raw))
	}
	return
}
