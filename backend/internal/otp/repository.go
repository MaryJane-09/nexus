package otp

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OTPRepository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *OTPRepository {
	return &OTPRepository{
		pool: pool,
	}
}

func (r *OTPRepository) Create(otp OTP) error {
	ctx := context.Background()

	_, err := r.pool.Exec(ctx, `INSERT INTO otps (email, code, expires_at, verified)
	VALUES ($1, $2, $3, $4) ON CONFLICT (email) DO UPDATE
	SET code = $2, expires_at = $3, verified = $4;`, otp.Email, otp.Code, otp.ExpiresAt, otp.Verified)
	if err != nil {
		return err
	}
	return nil
}

func (r *OTPRepository) FindByEmail(email string) (OTP, error) {

	ctx := context.Background()
	row := r.pool.QueryRow(ctx, `SELECT email, code, expires_at, verified FROM otps WHERE email = $1`, email)
	var otp OTP
	err := row.Scan(&otp.Email, &otp.Code, &otp.ExpiresAt, &otp.Verified)
	if err != nil {
		return OTP{}, errors.New("OTP not found")
	}
	return otp, nil

}

func (r *OTPRepository) Delete(email string) error {
	ctx := context.Background()

	_, err := r.pool.Exec(ctx,`DELETE FROM otps WHERE email = $1`, email)
	if err != nil {
		return err
	}
	return nil
}
