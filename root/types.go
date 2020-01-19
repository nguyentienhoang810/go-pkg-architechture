package root

type User struct {
	ID      int
	Name    string
	Email string
	Password string
	Phone string
	Address string
}

type Product struct {
	ID int
	Name string
	IDType int
	Description string
	UnitPrice float32
	PromoPrice float32
	Image string
	Unit string
	New int
}