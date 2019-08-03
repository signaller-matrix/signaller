package memory

import (
	"github.com/signaller-matrix/signaller/internal"
)

type BySize []internal.Room

func (a BySize) Len() int           { return len(a) }
func (a BySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySize) Less(i, j int) bool { return len(a[i].Users()) > len(a[j].Users()) }
