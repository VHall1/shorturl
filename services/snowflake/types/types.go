package types

type Snowflake interface {
	Generate() int64
}
