package query

import sq "github.com/Masterminds/squirrel"

const (
	usersTable      = "users"
	usersTableAlias = usersTable + " u"
)

type UserQuery struct {
	sq.StatementBuilderType
}

func User() *UserQuery {
	return &UserQuery{StatementBuilderType: sq.StatementBuilder.PlaceholderFormat(sq.Dollar)}
}

func (query UserQuery) SelectAll() sq.SelectBuilder {
	return query.buildSelectQuery()
}

func (query UserQuery) SelectOneByID(id int) sq.SelectBuilder {
	return query.buildSelectQuery().
		Where(sq.Eq{"u.id": id}).
		Limit(1)
}

func (query UserQuery) SelectOneByEmail(email string) sq.SelectBuilder {
	return query.buildSelectQuery().
		Where(sq.Eq{"u.email": email}).
		Limit(1)
}

func (query UserQuery) buildSelectQuery() sq.SelectBuilder {
	return query.Select(
		"u.id",
		"u.auth_token_hash",
		"u.age",
		"u.email",
		"u.password",
		"u.name",
		"u.surname",
		"u.sex",
		"u.city",
		"u.interests",
		"group_concat(users_friends.target_id) as friend_ids",
	).From(usersTableAlias).
		LeftJoin("users_friends ON users_friends.source_id = u.id").
		GroupBy("u.id")
}
