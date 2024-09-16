package request

import (
	"encoding/gob"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/auth/settings"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
	"github.com/gorilla/sessions"
)

type (
	// handles authentication state, keeps info about permanent login and
	// status of the MFA
	authUser struct {
		User *types.User

		// is user logged-in with "remember-me"?
		PermSession bool

		MFAStatus map[authType]authStatus
	}

	authStatus uint
	authType   string
)

const (
	// not enforced by user or by global settings
	authStatusDisabled authStatus = iota

	// when mfa is unconfigured on user but enforced globally
	// this can happen on when new user registers or when global enforcement is later enabled
	authStatusUnconfigured

	// enforced by user or global settings, but not yet verified
	authStatusPending

	// verified
	authStatusOK
)

const (
	// password does not really have any function,
	// more for demonstration purposes
	authByPassword = "password"
	authByEmailOTP = "email-otp"
	authByTOTP     = "totp"
)

func init() {
	gob.Register(&authUser{})
}

func NewAuthUser(s *settings.Settings, u *types.User, perm bool) *authUser {
	au := &authUser{
		User:        u,
		PermSession: perm,
		MFAStatus:   make(map[authType]authStatus),
	}

	au.set(s, u)
	return au
}

func (au *authUser) Update(s *settings.Settings, u *types.User) {
	au.set(s, u)
}

func (au *authUser) set(s *settings.Settings, u *types.User) {
	if u == nil || s == nil {
		return
	}

	if u.Meta == nil {
		u.Meta = &types.UserMeta{}
	}

	// User's MFA security policy
	umsp := u.Meta.SecurityPolicy.MFA

	// Global MFA security policy
	gmsp := s.MultiFactor

	mfaStatus := map[authType]authStatus{
		authByPassword: authStatusOK,
		authByEmailOTP: authStatusDisabled,
		authByTOTP:     authStatusDisabled,
	}

	// determinate mfa status for email OTP
	if !gmsp.EmailOTP.Enabled {
		mfaStatus[authByEmailOTP] = authStatusDisabled
	} else if umsp.EnforcedEmailOTP || gmsp.EmailOTP.Enforced {
		mfaStatus[authByEmailOTP] = authStatusPending
	}

	// determinate mfa status for TOTP
	if !gmsp.TOTP.Enabled {
		mfaStatus[authByTOTP] = authStatusDisabled
	} else if !umsp.EnforcedTOTP && gmsp.TOTP.Enforced {
		// TOTP not enforced on user but enforced globally
		mfaStatus[authByTOTP] = authStatusUnconfigured
	} else if gmsp.TOTP.Enforced {
		mfaStatus[authByTOTP] = authStatusPending
	}

	au.MFAStatus = mfaStatus
}

func (au authUser) DisabledEmailOTP() bool {
	return au.MFAStatus[authByEmailOTP] == authStatusDisabled
}

func (au authUser) PendingEmailOTP() bool {
	return au.MFAStatus[authByEmailOTP] == authStatusPending
}

func (au authUser) DisabledTOTP() bool {
	return au.MFAStatus[authByTOTP] == authStatusDisabled
}

func (au authUser) UnconfiguredTOTP() bool {
	return au.MFAStatus[authByTOTP] == authStatusUnconfigured
}

func (au authUser) PendingTOTP() bool {
	return au.MFAStatus[authByTOTP] == authStatusPending
}

// PendingMFA Returns true if any of MFAs are pending
func (au authUser) PendingMFA() bool {
	for _, st := range au.MFAStatus {
		if st == authStatusPending {
			return true
		}
	}

	return false
}

func (au *authUser) CompleteEmailOTP() {
	au.MFAStatus[authByEmailOTP] = authStatusOK
}

func (au *authUser) CompleteTOTP() {
	au.MFAStatus[authByTOTP] = authStatusOK
}

func (au *authUser) ResetTOTP() {
	au.MFAStatus[authByTOTP] = authStatusUnconfigured
}

func (au *authUser) Forget(ses *sessions.Session) {
	delete(ses.Values, keyAuthUser)
	delete(ses.Values, keyRememberMe)
}

func (au *authUser) Save(ses *sessions.Session) {
	ses.Values[keyAuthUser] = au

	// explicitly save roles
	ses.Values[keyRoles] = au.User.Roles()

	if au.PermSession {
		ses.Values[keyRememberMe] = true
	} else {
		delete(ses.Values, keyRememberMe)
	}
}
