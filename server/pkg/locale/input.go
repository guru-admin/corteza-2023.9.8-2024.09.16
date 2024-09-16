package locale

import (
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/xss"
)

func SanitizeMessage(in string) string {
	return xss.RichText(in)
}
