package queries

const (
	AddNewQuote = `INSERT INTO quotes (author, quote) VALUES ($1, $2) RETURNING id;`

	GetAllQuotesGrouped = `SELECT id, author, quote FROM quotes;`

	GetRandomQuoteGrouped = `SELECT id, author, quote FROM quotes ORDER BY RANDOM() LIMIT 1`

	GetByAuthorGrouped = `SELECT id, author, quote FROM quotes WHERE author = $1`

	DeleteById = `DELETE FROM quotes WHERE id = $1`
)
