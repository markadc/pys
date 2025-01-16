package py

type S = map[string]string

type A = map[string]interface{}

type RequestConfigs struct {
	Params S
	Header S
	Proxy  string
}
