package utils

import "database/sql"

func ParseNullString(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}
