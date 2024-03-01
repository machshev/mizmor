package scale

type Note uint8
type Offsets []int8

type Scale interface {
	fixed(id uint8) (Note, error)
	relative(fixedId uint8, relId uint8) ([]Note, error)
}
