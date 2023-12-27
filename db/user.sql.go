// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: user.sql

package db

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgtype"
)

const addShop = `-- name: AddShop :exec
INSERT INTO "shop" (
        "seller_name",
        "image_id",
        "name",
        "description",
        "enabled"
    )
VALUES(
        $1,
        NULL,
        $2,
        '',
        FALSE
    )
`

type AddShopParams struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	Name       string `form:"name" json:"name"`
}

func (q *Queries) AddShop(ctx context.Context, arg AddShopParams) error {
	_, err := q.db.Exec(ctx, addShop, arg.SellerName, arg.Name)
	return err
}

const addUser = `-- name: AddUser :exec
INSERT INTO "user" (
        "username",
        "password",
        "name",
        "email",
        "address",
        "role",
        "credit_card",
        "enabled"
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        'customer',
        '{}',
        TRUE
    )
`

type AddUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Address  string `form:"address" json:"address"`
}

func (q *Queries) AddUser(ctx context.Context, arg AddUserParams) error {
	_, err := q.db.Exec(ctx, addUser,
		arg.Username,
		arg.Password,
		arg.Name,
		arg.Email,
		arg.Address,
	)
	return err
}

const deleteRefreshToken = `-- name: DeleteRefreshToken :exec
UPDATE "user"
SET "refresh_token" = NULL
WHERE "refresh_token" = $1
`

func (q *Queries) DeleteRefreshToken(ctx context.Context, refreshToken string) error {
	_, err := q.db.Exec(ctx, deleteRefreshToken, refreshToken)
	return err
}

const findUserByRefreshToken = `-- name: FindUserByRefreshToken :one
SELECT "username",
    "role"
FROM "user"
WHERE "refresh_token" = $1
    AND "refresh_token_expire_date" > NOW()
`

type FindUserByRefreshTokenRow struct {
	Username string   `json:"username"`
	Role     RoleType `json:"role"`
}

func (q *Queries) FindUserByRefreshToken(ctx context.Context, refreshToken string) (FindUserByRefreshTokenRow, error) {
	row := q.db.QueryRow(ctx, findUserByRefreshToken, refreshToken)
	var i FindUserByRefreshTokenRow
	err := row.Scan(&i.Username, &i.Role)
	return i, err
}

const findUserInfoAndPassword = `-- name: FindUserInfoAndPassword :one
SELECT "username",
    "role",
    "password"
FROM "user"
WHERE "username" = $1
    OR "email" = $1
`

type FindUserInfoAndPasswordRow struct {
	Username string   `json:"username"`
	Role     RoleType `json:"role"`
	Password string   `json:"password"`
}

// user can enter both username and email to verify
// but writing "usernameOrEmail" is too long
func (q *Queries) FindUserInfoAndPassword(ctx context.Context, username string) (FindUserInfoAndPasswordRow, error) {
	row := q.db.QueryRow(ctx, findUserInfoAndPassword, username)
	var i FindUserInfoAndPasswordRow
	err := row.Scan(&i.Username, &i.Role, &i.Password)
	return i, err
}

const setRefreshToken = `-- name: SetRefreshToken :exec
UPDATE "user"
SET "refresh_token" = $1,
    "refresh_token_expire_date" = $2
WHERE "username" = $3
`

type SetRefreshTokenParams struct {
	RefreshToken string             `json:"refresh_token"`
	ExpireDate   pgtype.Timestamptz `json:"expire_date"`
	Username     string             `json:"username"`
}

func (q *Queries) SetRefreshToken(ctx context.Context, arg SetRefreshTokenParams) error {
	_, err := q.db.Exec(ctx, setRefreshToken, arg.RefreshToken, arg.ExpireDate, arg.Username)
	return err
}

const userExists = `-- name: UserExists :one
SELECT EXISTS (
        SELECT 1
        FROM "user"
        WHERE "username" = $1
            OR "email" = $2
    )
`

type UserExistsParams struct {
	Username string `json:"username"`
	Email    string `form:"email" json:"email"`
}

func (q *Queries) UserExists(ctx context.Context, arg UserExistsParams) (bool, error) {
	row := q.db.QueryRow(ctx, userExists, arg.Username, arg.Email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const userGetCreditCard = `-- name: UserGetCreditCard :one
SELECT "credit_card"
FROM "user"
WHERE "username" = $1
`

func (q *Queries) UserGetCreditCard(ctx context.Context, username string) (json.RawMessage, error) {
	row := q.db.QueryRow(ctx, userGetCreditCard, username)
	var credit_card json.RawMessage
	err := row.Scan(&credit_card)
	return credit_card, err
}

const userGetInfo = `-- name: UserGetInfo :one
SELECT "name",
    "email",
    "address",
    "image_id" as "image_url"
FROM "user" u
WHERE u."username" = $1
`

type UserGetInfoRow struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Address  string `form:"address" json:"address"`
	ImageUrl string `json:"image_url" swaggertype:"string"`
}

func (q *Queries) UserGetInfo(ctx context.Context, username string) (UserGetInfoRow, error) {
	row := q.db.QueryRow(ctx, userGetInfo, username)
	var i UserGetInfoRow
	err := row.Scan(
		&i.Name,
		&i.Email,
		&i.Address,
		&i.ImageUrl,
	)
	return i, err
}

const userGetPassword = `-- name: UserGetPassword :one
SELECT "password"
FROM "user"
WHERE "username" = $1
`

func (q *Queries) UserGetPassword(ctx context.Context, username string) (string, error) {
	row := q.db.QueryRow(ctx, userGetPassword, username)
	var password string
	err := row.Scan(&password)
	return password, err
}

const userUpdateCreditCard = `-- name: UserUpdateCreditCard :one
UPDATE "user"
SET "credit_card" = $2
WHERE "username" = $1
RETURNING "credit_card"
`

type UserUpdateCreditCardParams struct {
	Username   string          `json:"username"`
	CreditCard json.RawMessage `json:"credit_card"`
}

func (q *Queries) UserUpdateCreditCard(ctx context.Context, arg UserUpdateCreditCardParams) (json.RawMessage, error) {
	row := q.db.QueryRow(ctx, userUpdateCreditCard, arg.Username, arg.CreditCard)
	var credit_card json.RawMessage
	err := row.Scan(&credit_card)
	return credit_card, err
}

const userUpdateInfo = `-- name: UserUpdateInfo :one
UPDATE "user"
SET "name" = COALESCE($2, "name"),
    "email" = COALESCE($3, "email"),
    "address" = COALESCE($4, "address"),
    "image_id" = CASE
        WHEN $5::TEXT = '' THEN "image_id"
        ELSE $5::TEXT
    END
WHERE "username" = $1
RETURNING "name",
    "email",
    "address",
    "image_id" as "image_url"
`

type UserUpdateInfoParams struct {
	Username string `json:"username"`
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Address  string `form:"address" json:"address"`
	ImageID  string `json:"image_id"`
}

type UserUpdateInfoRow struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Address  string `form:"address" json:"address"`
	ImageUrl string `json:"image_url" swaggertype:"string"`
}

func (q *Queries) UserUpdateInfo(ctx context.Context, arg UserUpdateInfoParams) (UserUpdateInfoRow, error) {
	row := q.db.QueryRow(ctx, userUpdateInfo,
		arg.Username,
		arg.Name,
		arg.Email,
		arg.Address,
		arg.ImageID,
	)
	var i UserUpdateInfoRow
	err := row.Scan(
		&i.Name,
		&i.Email,
		&i.Address,
		&i.ImageUrl,
	)
	return i, err
}

const userUpdatePassword = `-- name: UserUpdatePassword :one
UPDATE "user"
SET "password" = $2
WHERE "username" = $1
RETURNING "name",
    "email",
    "address",
    "image_id" as "image_url"
`

type UserUpdatePasswordParams struct {
	Username    string `json:"username"`
	NewPassword string `json:"new_password"`
}

type UserUpdatePasswordRow struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Address  string `form:"address" json:"address"`
	ImageUrl string `json:"image_url" swaggertype:"string"`
}

func (q *Queries) UserUpdatePassword(ctx context.Context, arg UserUpdatePasswordParams) (UserUpdatePasswordRow, error) {
	row := q.db.QueryRow(ctx, userUpdatePassword, arg.Username, arg.NewPassword)
	var i UserUpdatePasswordRow
	err := row.Scan(
		&i.Name,
		&i.Email,
		&i.Address,
		&i.ImageUrl,
	)
	return i, err
}
