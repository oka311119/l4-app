package date

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatAsISO8601(t *testing.T) {
	// 特定の日時を作成
	tm := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)

	// ISO 8601形式の文字列に変換
	s := FormatAsISO8601(tm)

	assert.Equal(t, "2022-01-01T00:00:00Z", s)
}