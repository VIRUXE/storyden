// Code generated by enumerator. DO NOT EDIT.

package datagraph

import (
	"database/sql/driver"
	"fmt"
)

type Kind struct {
	v kindEnum
}

var (
	KindPost = Kind{kindPost}
	KindNode = Kind{kindNode}
)

func (r Kind) Format(f fmt.State, verb rune) {
	switch verb {
	case 's':
		fmt.Fprint(f, r.v)
	case 'q':
		fmt.Fprintf(f, "%q", r.String())
	default:
		fmt.Fprint(f, r.v)
	}
}
func (r Kind) String() string {
	return string(r.v)
}
func (r Kind) MarshalText() ([]byte, error) {
	return []byte(r.v), nil
}
func (r *Kind) UnmarshalText(__iNpUt__ []byte) error {
	s, err := NewKind(string(__iNpUt__))
	if err != nil {
		return err
	}
	*r = s
	return nil
}
func (r Kind) Value() (driver.Value, error) {
	return r.v, nil
}
func (r *Kind) Scan(__iNpUt__ any) error {
	s, err := NewKind(fmt.Sprint(__iNpUt__))
	if err != nil {
		return err
	}
	*r = s
	return nil
}
func NewKind(__iNpUt__ string) (Kind, error) {
	switch __iNpUt__ {
	case string(kindPost):
		return KindPost, nil
	case string(kindNode):
		return KindNode, nil
	default:
		return Kind{}, fmt.Errorf("invalid value for type 'Kind': '%s'", __iNpUt__)
	}
}
