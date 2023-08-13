package repository

const (
	queryInsertNewUser = `
		INSERT INTO users(
			full_name, phone_number, password
		)
		VALUES(
			$1, $2, $3
		)
	`
	queryGetUserByPhoneNumber = `
		SELECT
			id, password
		FROM users
		WHERE phone_number = $1
	`
)
