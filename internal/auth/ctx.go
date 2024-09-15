package auth

import "context"

type ctxKey int

const authKey ctxKey = 0

func WithSubject(ctx context.Context, subject string) context.Context {
	return context.WithValue(ctx, authKey, subject)
}

func GetSubject(ctx context.Context) string {
	subject, _ := ctx.Value(authKey).(string)
	return subject
}
