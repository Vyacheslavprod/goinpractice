// Тестирование с помощью заглушки
package library_test

import "testing"

type MockMessage struct {
	email, subject string
	body []byte
}

func (m *MockMessage) Send(email, subject string, body []byte) error {
	m.email = email
	m.subject = subject
	m.body = body
	return nil
}

func TestAlert(t *testing.T) {
	msgr := new(MockMessage) // Создание нового MockMessage
	body := []byte("Critical Error")

	Alert(msgr, body) // Вызов метода Alert заглушки

	if msgr.subject != "Cretical Error" {
		t.Errorf("Expected 'Critical Error', Got '%s'", msgr.subject)
	}
}