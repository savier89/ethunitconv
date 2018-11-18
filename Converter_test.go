package ethunitconv

import (
	"fmt"
	"testing"
)

func TestConverter(t *testing.T) {
	ether := FromWei("418160973408927260000000000", "Ether")
	fmt.Println(ether)
	ether2wei := ToWei(ether, "Ether")
	fmt.Println(ether2wei)
}
