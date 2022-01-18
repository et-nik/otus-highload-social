package api

import "context"

type contextKey int

const sessionKey contextKey = iota

func sessionFromContext(ctx context.Context) *Session {
	session, _ := ctx.Value(sessionKey).(*Session)

	return session
}

func contextWithSession(ctx context.Context, session *Session) context.Context {
	return context.WithValue(ctx, sessionKey, session)
}
