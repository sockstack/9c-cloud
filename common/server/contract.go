package server

type serverInterface interface {
	Run(address string)
}

type serverDriver interface {
	ServiceRegister(s interface{}) error
	Address(address string) error
}
