package scale

type Note uint8
type Offsets []int8

type Scale interface {
	Fixed(id uint8) (Note, error)
	Relative(fixedId uint8, relId uint8) ([]Note, error)
}
