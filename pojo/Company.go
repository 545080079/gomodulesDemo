package pojo

type Company struct {
	talker Talker
}

func NewCompany(t Talker) *Company {
	return &Company{
		talker: t,
	}
}

func (c *Company) Meeting(guestName, guestRole string) string {
	return c.talker.SayHello(guestName, guestRole)
}
