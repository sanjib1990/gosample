package app

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/fatih/structs"
	"golang.org/x/exp/slices"
	"reflect"
)

var (
	StringEqComparator = func(a string) func(s string) bool {
		return func(s string) bool {
			return a == s
		}
	}
)

type List[T interface{}] struct {
	items []T
}

func (l *List[T]) IsEmpty() bool {
	return l.items == nil || len(l.items) == 0
}

func (l *List[T]) Sort(fn func(T, T) bool) *List[T] {
	slices.SortStableFunc(l.items, fn)
	return l
}

// Contains item in the slice. fn should return 0 if a == b,
// a negative number if a < b and a positive number if a > b
// The slice must be sorted in increasing order
func (l *List[T]) Contains(fn func(T) bool) bool {
	return slices.IndexFunc(l.items, fn) >= 0
}

func (l *List[T]) GetIndex(f func(T) bool) int {
	return slices.IndexFunc(l.items, f)
}

func (l *List[T]) Length() int {
	return len(l.items)
}

func (l *List[T]) SetItems(items []T) *List[T] {
	if items == nil {
		return l
	}

	l.items = items

	return l
}

func (l *List[T]) Append(item T) *List[T] {
	l.items = append(l.items, item)

	return l
}

func (l *List[T]) GetItems() []T {
	return l.items
}

type ISanitizer[T ISanitisable] interface {
	SetSubject(item T) ISanitizer[T]
	GetSanitizedItem() T
	SetFields(fields []string) ISanitizer[T]
}

type ISanitisable interface {
	GetXssSanitizeFields() []string
}

type Sanitizer[T ISanitisable] struct {
	item   T
	fields *List[string]
}

func (s *Sanitizer[T]) SetSubject(item T) ISanitizer[T] {
	s.item = item

	s.SetFields(item.GetXssSanitizeFields())

	return s
}

func (s *Sanitizer[T]) GetSanitizedItem() T {
	fields := structs.Fields(s.item)
	_ = reflect.ValueOf(s.item)

	//if !s.fields.Contains(StringEqComparator(fieldName)) {
	//	continue
	//}
	for _, field := range fields {
		_ = field.Tag("json")
		switch field.Kind().String() {
		case "string":
		}
	}

	return s.item
}

func (s *Sanitizer[T]) SetFields(fields []string) ISanitizer[T] {
	s.fields = &List[string]{}
	s.fields.SetItems(fields)
	return s
}

type Profile struct {
	Image   string   `json:"image"`
	Address *Address `json:"address"`
	Hobbies []string `json:"hobbies"`
}

func (p *Profile) GetXssSanitizeFields() []string {
	return []string{
		"image", "address",
	}
}

type User struct {
	Name    string   `json:"name,omitempty,-"`
	Age     uint64   `json:"age,omitempty,-"`
	Profile *Profile `json:"profile,omitempty,-"`
}

func (u *User) GetXssSanitizeFields() []string {
	return []string{"name", "age", "profile"}
}

func main1() {
	profile := &Profile{
		Image: gofakeit.ImageURL(300, 300),
		Address: &Address{
			StreetName: gofakeit.Address().Address,
		},
	}

	user := &User{
		Name:    gofakeit.Name(),
		Age:     uint64(gofakeit.Number(10, 60)),
		Profile: profile,
	}

	stn := Sanitizer[*User]{}
	stn.SetSubject(user)
	user = stn.GetSanitizedItem()
}
