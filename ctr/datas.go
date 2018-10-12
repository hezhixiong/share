package ctr

const (
	PAGE_DEFAULT      = 1
	PAGE_SIZE_DEFAULT = 20
	PAGE_SIZE_MIN     = 10
	PAGE_SIZE_MAX     = 100
)

var codeMsg = map[int]string{
	-1: "error",
	0:  "success",
}
