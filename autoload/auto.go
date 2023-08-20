package autoload

import "github.com/HuntClauss/dotenv"

func init() {
	panic(dotenv.LoadEnv(".env"))
}
