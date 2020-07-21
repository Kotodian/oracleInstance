package pager

type Pager struct {
	Page   int         `json:"page"`
	Size   int         `json:"size"`
	Offset int         `json:"-"`
	Total  int         `json:"total"`
	Data   interface{} `json:"data"`
	Next   bool        `json:"next"`
	Prev   bool        `json:"prev"`
}

func NewPager(page int, size int) *Pager {
	pager := &Pager{
		Page: page,
		Size: size,
		Data: nil,
		Next: false,
		Prev: false,
	}
	if page > 1 {
		pager.Prev = true
	}
	pager.Offset = (pager.Page - 1) * pager.Size

	return pager
}

func (self *Pager) SetTotal(total int) {
	self.Total = total

	if self.Offset+self.Size < self.Total {
		self.Next = true
	}
}
