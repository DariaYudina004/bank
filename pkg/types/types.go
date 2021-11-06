package types

// Money  представляет собой денежную сумму в минимальных единицах (центы , копейки , дирамы и т.д.)
type Money int64

//Currency представляют код валюты
type Currency string

//Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

//PAN  представляет номер карты
type PAN string

//Card представляет информацию о платежной карте
type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}

type Category string //категории в которых был совершен платеж(авто , рестораны,аптеки итд.)

//Status представляет собой статус платежа
type Status string

//Предопределенные статусы платежей
const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

type Payment struct { // информация о платeже
	ID       int
	Amount   Money
	Category Category
	Status   Status
}

