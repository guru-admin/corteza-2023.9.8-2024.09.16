package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/auth/request"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/auth/settings"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
	oauth2models "github.com/go-oauth2/oauth2/v4/models"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func Test_oauth2AuthorizeSuccess(t *testing.T) {
	var (
		user = makeMockUser()

		req = &http.Request{
			Form:     url.Values{},
			PostForm: url.Values{},
		}

		oauthService oauth2Service
		authService  authService
		authHandlers *AuthHandlers
		authReq      *request.AuthReq
	)

	tcc := []testingExpect{
		{
			name:     "authorization success",
			payload:  -1,
			err:      "",
			template: "",
			fn: func(_ *settings.Settings) {
				oauthService = &oauth2ServiceMocked{
					handleAuthorizeRequest: func(w http.ResponseWriter, r *http.Request) error {
						return nil
					},
				}
			},
		},
		{
			name:     "authorization failure",
			payload:  http.StatusInternalServerError,
			err:      "not authorized",
			template: TmplInternalError,
			fn: func(_ *settings.Settings) {
				oauthService = &oauth2ServiceMocked{
					handleAuthorizeRequest: func(w http.ResponseWriter, r *http.Request) error {
						return fmt.Errorf("not authorized")
					},
				}
			},
		},
	}

	for _, tc := range tcc {
		t.Run(tc.name, func(t *testing.T) {
			rq := require.New(t)

			authSettings := &settings.Settings{}

			tc.fn(authSettings)

			authHandlers = &AuthHandlers{
				Log:         zap.NewNop(),
				AuthService: authService,
				Settings:    authSettings,
				OAuth2:      oauthService,
			}
			authReq = prepareClientAuthReq(authHandlers, req, user)

			err := authHandlers.oauth2Authorize(authReq)

			rq.NoError(err)
			rq.Equal(tc.template, authReq.Template)
			rq.Equal(tc.payload, authReq.Status)

			if tc.err != "" {
				rq.EqualError(errors.New(tc.err), authReq.Data["error"].(error).Error())
			}
		})
	}
}

func Test_oauth2AuthorizeSuccessSetParams(t *testing.T) {
	var (
		user = makeMockUser()

		req = &http.Request{
			Form:     url.Values{},
			PostForm: url.Values{},
		}

		authService  authService
		authHandlers *AuthHandlers
		authReq      *request.AuthReq

		authSettings = &settings.Settings{}

		rq = require.New(t)
	)

	oauthService := &oauth2ServiceMocked{
		handleAuthorizeRequest: func(w http.ResponseWriter, r *http.Request) error {
			return nil
		},
	}

	authHandlers = &AuthHandlers{
		Log:         zap.NewNop(),
		AuthService: authService,
		Settings:    authSettings,
		OAuth2:      oauthService,
	}
	authReq = prepareClientAuthReq(authHandlers, req, user)
	authReq.Session.Values["oauth2AuthParams"] = url.Values{"foo": []string{"bar"}}

	err := authHandlers.oauth2Authorize(authReq)

	rq.NoError(err)
	rq.Equal("", authReq.Template)
	rq.Equal(-1, authReq.Status)
	rq.Equal(nil, authReq.Data["error"])
}

func Test_generateIdToken(t *testing.T) {
	var (
		req = require.New(t)
	)

	signed, err := generateIdToken(
		&types.User{},
		&types.AuthClient{
			Secret: "correct horse battery stable",
		},
		&oauth2models.Token{},
		"http://cortezaproject.org",
	)

	req.NoError(err)
	req.NotEmpty(signed)
}

func Test_SubSplitRoles(t *testing.T) {
	type (
		exp struct {
			id string
			i  string
			ii []string
		}
	)
	var (
		req = require.New(t)
		ti  = &oauth2models.Token{}
		d   = make(map[string]interface{})

		tii = []exp{
			{id: "1", i: "1", ii: []string{}},
			{id: "1 2", i: "1", ii: []string{"2"}},
			{id: "1 2 33 444", i: "1", ii: []string{"2", "33", "444"}},
		}
	)

	for _, v := range tii {
		ti.SetUserID(v.id)
		SubSplit(ti, d)

		req.Equal(v.i, d["sub"])

		if _, is := d["roles"]; is {
			req.Equal(v.ii, d["roles"])
		}
	}
}
