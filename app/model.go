package app

type Const struct {
	Port string `env:"PORT"`
}

// HTTPMessgae json message
type HTTPMessgae struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
