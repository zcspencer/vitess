package tableacl

import "fmt"

var allAccess SimpleACL

const (
	ALL = "*"
)

func init() {
	allAccess = SimpleACL(map[string]bool{ALL: true})
}

// NewACL returns an ACL with the specified entries
func NewACL(entries []string) (ACL, error) {
	a := SimpleACL(map[string]bool{})
	for _, e := range entries {
		a[e] = true
	}
	return a, nil
}

// SimpleACL keeps all entries in a unique in-memory list
type SimpleACL map[string]bool

// Check checks the membership of a principal in this ACL
func (a SimpleACL) Check(principal string) error {
	if a[principal] || a[ALL] {
		return nil
	}
	return fmt.Errorf("%v not found in list", principal)
}
