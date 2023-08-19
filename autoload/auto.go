package autoload

import "dotenv/env"

func init() {
	panic(env.LoadEnv(".env"))
}
