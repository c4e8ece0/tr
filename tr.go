// --------------------------------------------------------------------------
// Сокращение строки в int (положительное целое)
// --------------------------------------------------------------------------

package tr

// Основная структура
type tr struct {
	cnt     int            // счётчик термов
	str_def string         // строка по умолчанию для пустых
	has_err bool           // флаг ошибки на предыдущем поиске
	s2i     map[string]int // перевод строки в идентификатор
	i2s     map[int]string // перевод идентификатора в строку
}

// Создание объекта
func NewTr() *tr {
	s := &tr{}
	s.s2i = make(map[string]int)
	s.i2s = make(map[int]string)
	s.cnt = 0
	s.str_def = "[not_found]"
	return s
}

// Перевод строки в идентификатор
func (s *tr) Id(str string) int {
	r, ok := s.s2i[str]
	if !ok {
		s.s2i[str] = s.cnt
		s.i2s[s.cnt] = str
		r = s.cnt
		s.cnt++
	}
	return r
}

// Проверка существования строки
func (s *tr) StrExists(str string) bool {
	_, ok := s.s2i[str]
	return !ok
}

// Перевод идентификатора в строку
func (s *tr) Str(id int) string {
	r, ok := s.i2s[id]
	s.has_err = false
	if !ok {
		s.has_err = true
		r = s.str_def
	}
	return r
}

// Проверка была ли ошибка
func (s *tr) HasErr() bool {
	return s.has_err
}

// Дефолтная строка для перевода id=>str
func (s *tr) SetEmpty(str string) *tr {
	s.str_def = str
	return s
}
