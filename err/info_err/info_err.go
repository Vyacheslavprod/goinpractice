// Возвращает дополнительную информацию о местоположении ошибки
package info_err

import "fmt"

type ParseError struct {
	Message string // Сообщение об ошибке
	Line, Char int // Информация о местоположении
}

// Реализация интерфейса Error
func (p *ParseError) Error() string {
	format := "%s on Line %d, Char %d"
	return fmt.Sprintf(format, p.Message, p.Line, p.Char)
}