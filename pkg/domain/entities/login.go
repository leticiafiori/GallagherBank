package entities

type Login struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}
