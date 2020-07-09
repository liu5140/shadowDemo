package security

//IPasswordEncoder encoder password
type IPasswordEncoder interface {
	Encode(password string) string
	Matches(password string, encodedPassword string) bool
}

type FPasswordEncoderFactory func() IPasswordEncoder

var PasswordEncoderFactories = make(map[string]FPasswordEncoderFactory)

func RegisterPasswordEncoder(name string, factory FPasswordEncoderFactory) {
	PasswordEncoderFactories[name] = factory
}

func PasswordEncoderInstance(name string) IPasswordEncoder {
	factory := PasswordEncoderFactories[name]
	return factory()
}
