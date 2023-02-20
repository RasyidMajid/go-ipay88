package go_ipay88

type Env int8

const (
	_ Env = iota

	Sandbox

	Production
)

var typeString = map[Env]string{
	Sandbox:    "https://sandbox.ipay88.co.id",
	Production: "",
}

func (e Env) String() string {
	for k, v := range typeString {
		if k == e {
			return v
		}
	}

	return "undefined"
}
