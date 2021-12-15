package go_dic

type inverted_index struct {
	sid int
	// time string
	position int
}

func NewInverted_index(sid int, p int) *inverted_index {
	return &inverted_index{
		sid:      sid,
		position: p,
	}
}
