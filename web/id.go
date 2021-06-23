package web

import "fmt"

type WebId uint32

type WebIds []WebId

const (
	webPrefix = "web"
	InvalidId = 0xffffffff
)

func (id WebId) Valid() bool {
	return id != InvalidId
}

func ToWebId(s string) (id WebId) {
	n, err := fmt.Sscanf(s, webPrefix+"%d", &id)
	if n < 1 || err != nil {
		id = InvalidId
	}
	return
}

func (id WebId) String() string {
	return fmt.Sprintf("%s%d", webPrefix, id)
}

func (id WebId) Index(i int) string {
	return fmt.Sprintf("%s%d-%d", webPrefix, id, i)
}

func (ids WebIds) Len() int {
	return len(ids)
}

func (ids WebIds) Less(i, j int) bool {
	return ids[i] < ids[j]
}

func (ids WebIds) Swap(i, j int) {
	ids[i], ids[j] = ids[j], ids[i]
}
