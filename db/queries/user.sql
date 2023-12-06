-- name: UserGetInfo :one
SELECT "name",
    "email",
    "image_id",
    "enabled"
FROM "user" u
WHERE u."username" = $1;

-- name: UserUpdateInfo :one
UPDATE "user"
SET "name" = COALESCE($2, "name"),
    "email" = COALESCE($3, "email"),
    "address" = COALESCE($4, "address"),
    "image_id" = COALESCE($5, "image_id")
WHERE "username" = $1
RETURNING "name",
    "email",
    "image_id",
    "enabled";

-- name: UserGetPassword :one
SELECT "password"
FROM "user"
WHERE "username" = $1;

-- name: UserUpdatePassword :one
UPDATE "user"
SET "password" = sqlc.arg(new_password)
WHERE "username" = $1
RETURNING "name",
    "email",
    "address",
    "image_id",
    "enabled";

-- name: UserGetCreditCard :one
SELECT "credit_card"
FROM "user"
WHERE "username" = $1;

-- name: UserUpdateCreditCard :one
UPDATE "user"
SET "credit_card" = $2
WHERE "username" = $1
RETURNING "credit_card";

-- name: AddUser :exec
INSERT INTO "user" (
        "username",
        "password",
        "name",
        "email",
        "address",
        "role",
        "credit_card"
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        'customer',
        '{}'
    );

-- name: UserExists :one
SELECT EXISTS (
        SELECT 1
        FROM "user"
        WHERE "username" = $1
            OR "email" = $2
    );
