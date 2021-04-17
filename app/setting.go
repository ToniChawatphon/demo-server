package app

import "os"

var (
	Setting *Const
)

func Init() {
	Setting = &Const{}
	Setting.Port = os.Getenv("PORT")
}
