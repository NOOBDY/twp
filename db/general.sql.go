// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: general.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getProductInfo = `-- name: GetProductInfo :one

SELECT
    "id",
    "name",
    "description",
    "price",
    "image_id",
    "exp_date",
    "stock",
    "sales"
FROM "product"
WHERE
    "id" = $1
    AND "enabled" = TRUE
`

type GetProductInfoRow struct {
	ID          int32              `json:"id" param:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       pgtype.Numeric     `json:"price"`
	ImageID     pgtype.UUID        `json:"image_id"`
	ExpDate     pgtype.Timestamptz `json:"exp_date"`
	Stock       int32              `json:"stock"`
	Sales       int32              `json:"sales"`
}

func (q *Queries) GetProductInfo(ctx context.Context, id int32) (GetProductInfoRow, error) {
	row := q.db.QueryRow(ctx, getProductInfo, id)
	var i GetProductInfoRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.ImageID,
		&i.ExpDate,
		&i.Stock,
		&i.Sales,
	)
	return i, err
}

const getSellerNameByShopID = `-- name: GetSellerNameByShopID :one

SELECT "seller_name" FROM "shop" WHERE "id" = $1
`

func (q *Queries) GetSellerNameByShopID(ctx context.Context, id int32) (string, error) {
	row := q.db.QueryRow(ctx, getSellerNameByShopID, id)
	var seller_name string
	err := row.Scan(&seller_name)
	return seller_name, err
}

const getShopCoupons = `-- name: GetShopCoupons :many

SELECT
    "id",
    "type",
    "scope",
    "name",
    "description",
    "discount",
    "start_date",
    "expire_date"
FROM "coupon"
WHERE
    "shop_id" = $1
    OR "scope" = 'global'
ORDER BY "id" ASC
LIMIT $2
OFFSET $3
`

type GetShopCouponsParams struct {
	ShopID pgtype.Int4 `json:"-"`
	Limit  int32       `json:"limit"`
	Offset int32       `json:"offset"`
}

type GetShopCouponsRow struct {
	ID          int32              `json:"-" param:"id"`
	Type        CouponType         `json:"type"`
	Scope       CouponScope        `json:"scope"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Discount    pgtype.Numeric     `json:"discount" swaggertype:"number"`
	StartDate   pgtype.Timestamptz `json:"start_date" swaggertype:"string"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
}

func (q *Queries) GetShopCoupons(ctx context.Context, arg GetShopCouponsParams) ([]GetShopCouponsRow, error) {
	rows, err := q.db.Query(ctx, getShopCoupons, arg.ShopID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetShopCouponsRow{}
	for rows.Next() {
		var i GetShopCouponsRow
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Scope,
			&i.Name,
			&i.Description,
			&i.Discount,
			&i.StartDate,
			&i.ExpireDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getShopInfo = `-- name: GetShopInfo :one

SELECT
    "seller_name",
    "image_id",
    "name",
    "description"
FROM "shop"
WHERE
    "seller_name" = $1
    AND "enabled" = TRUE
`

type GetShopInfoRow struct {
	SellerName  string      `json:"seller_name" param:"seller_name"`
	ImageID     pgtype.UUID `json:"image_id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
}

func (q *Queries) GetShopInfo(ctx context.Context, sellerName string) (GetShopInfoRow, error) {
	row := q.db.QueryRow(ctx, getShopInfo, sellerName)
	var i GetShopInfoRow
	err := row.Scan(
		&i.SellerName,
		&i.ImageID,
		&i.Name,
		&i.Description,
	)
	return i, err
}

const getTagInfo = `-- name: GetTagInfo :one

SELECT "id", "name" FROM "tag" WHERE "id" = $1
`

type GetTagInfoRow struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) GetTagInfo(ctx context.Context, id int32) (GetTagInfoRow, error) {
	row := q.db.QueryRow(ctx, getTagInfo, id)
	var i GetTagInfoRow
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const shopExists = `-- name: ShopExists :one

SELECT "id"
FROM "shop" AS s
WHERE
    s."seller_name" = $1
    AND s."enabled" = TRUE
`

func (q *Queries) ShopExists(ctx context.Context, sellerName string) (int32, error) {
	row := q.db.QueryRow(ctx, shopExists, sellerName)
	var id int32
	err := row.Scan(&id)
	return id, err
}
