package repository

import (
	"time"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func dateStringToPGDate(value string) (pgtype.Date, error) {
	t, err := time.Parse("2006-05-01", value)
	if err != nil {
		return pgtype.Date{}, err
	}
	return pgtype.Date{Time: t, Valid: true}, nil
}

func pgDateToString(d pgtype.Date) string {
	if !d.Valid {
		return ""
	}
	return d.Time.Format("2006-05-01")
}

func uuidStringToPgUUID(id string) (pgtype.UUID, error) {
	u, err := uuid.Parse(id)
	if err != nil {
		return pgtype.UUID{}, err
	}
	return pgtype.UUID{Bytes: u, Valid: true}, nil
}

func pgTypeToString(u pgtype.UUID) string {
	if !u.Valid {
		return ""
	}
	return uuid.UUID(u.Bytes).String()
}
