package main

import (
	"fmt"
	"net/http"
)

type Handler struct {
	// ...
}

var _ http.Handler = (*Handler)(nil)

func (h *Handler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	// ...
}

type LogHandler struct {
	h http.Handler
}

var _ http.Handler = LogHandler{}

func (h LogHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	// ...
}

type F interface {
	f1()
	f2()
}

type S1 struct{}

func (s S1) f1() {}
func (s S1) f2() {}

type S2 struct{}

func (s *S2) f1() {}
func (s S2) f2()  {}

var _ F = (*S1)(nil)
var _ F = &S1{}
var _ F = &S2{}

type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}

func (s *S) Write(str string) {
	s.data = str
}

func main() {
	s1Val := S1{}
	s1Ptr := &S1{}
	//s2Val := S2{}
	s2Ptr := &S2{}

	var i F
	i = s1Val
	i = s1Ptr
	i = s2Ptr
	// i = s2Val // cannot use s2Val (variable of type S2) as F value in assignment: S2 does not implement F (method f1 has pointer receiver)
	fmt.Println(i)

	ReciversAndInterface()
}

func ReciversAndInterface() {

	// We cannot get pointers to values stored in maps, because they are not
	// addressable values.
	sVals := map[int]S{1: {"A"}}

	// We can call Read on values stored in the map because Read
	// has a value receiver, which does not require the value to
	// be addressable.
	sVals[1].Read()

	// We cannot call Write on values stored in the map because Write
	// has a pointer receiver, and it's not possible to get a pointer
	// to a value stored in a map.
	//
	ss := sVals[1]
	ss.Write("test")
	//sVals[1].Write("ffff")

	sPtrs := map[int]*S{1: {"A"}}

	// You can call both Read and Write if the map stores pointers,
	// because pointers are intrinsically addressable.
	sPtrs[1].Read()
	sPtrs[1].Write("test")
}
