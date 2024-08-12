// Структура сообщения
package library_test

type Message struct {
	// ...
}

func (m *Message) Send(email, subject string, body []byte) error  {
	// ...
	return nil
}

// Использование интерфейса

// Определение интерфейса, описывающего метод, используемый типом Message
type Messager interface {
	Send(email, subject string, body []byte) error
}

// Передача интерфейса вместо типа Message
func Alert(m Messager, problem []byte) error {
	return m.Send("noc@example.com", "Critical Error", problem)
}