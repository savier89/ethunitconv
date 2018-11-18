package ethunitconv

import (
	"fmt"
	"math/big"
	"strings"
)

// Factor type for enum
type Factor int

const (
	// Wei 1 wei
	Wei Factor = 0
	// KWei (babbage) 1e3 wei
	KWei Factor = 3
	// MWei (lovelace) 1e6 wei
	MWei Factor = 6
	// GWei (shannon) 1e9 wei
	GWei Factor = 9
	// Szabo (microether) 1e12 wei
	Szabo Factor = 12
	// Finney (milliether) 1e15 wei
	Finney Factor = 15
	// Ether 1e18 wei
	Ether Factor = 18
	// KEther 1e21 wei
	KEther Factor = 21
	// MEther 1e24 wei
	MEther Factor = 24
	// GEther 1e27 wei
	GEther Factor = 27
)

// Unit Base struct
type Unit struct {
	Name      string
	WeiFactor *big.Int
}

// FromWei func calculate number / unitfactor return string
func FromWei(number string, unitname string) string {
	unit := new(Unit)
	unit = unit.InitUnit(unitname)
	bigFloatNumber, _ := new(big.Float).SetString(number)
	unitWeiFactor := new(big.Float).SetInt(unit.GetWeiFactor())
	tmp := new(big.Float)
	tmp.Quo(bigFloatNumber, unitWeiFactor)
	tmpstr := fmt.Sprintf("%.10f", tmp)

	return tmpstr
}

// ToWei func calculate number * unitfactor return string
func ToWei(number string, unitname string) string {
	unit := new(Unit)
	unit = unit.InitUnit(unitname)
	bigFloatNumber, _ := new(big.Float).SetString(number)
	unitWeiFactor := new(big.Float).SetInt(unit.GetWeiFactor())
	tmp := new(big.Float)
	tmp.Mul(bigFloatNumber, unitWeiFactor)
	tmpstr := fmt.Sprintf("%.f", tmp)

	return tmpstr
}

// InitUnit func param unitname(string) retun Unit struct
// by default Wei factor
func (unit *Unit) InitUnit(unitname string) *Unit {
	name := strings.ToLower(unitname)
	units := map[string]Factor{
		"wei":    Wei,
		"kwei":   KWei,
		"mwei":   MWei,
		"gwei":   GWei,
		"szabo":  Szabo,
		"finney": Finney,
		"ether":  Ether,
		"kether": KEther,
		"mether": MEther,
		"gether": GEther,
	}

	var unitfactor, exp = big.NewInt(int64(units[name])), big.NewInt(10)
	exp.Exp(exp, unitfactor, nil)
	unit.Name = name
	unit.WeiFactor = exp

	return unit
}

// GetWeiFactor func return unit.WeiFactor
func (unit *Unit) GetWeiFactor() *big.Int {
	return unit.WeiFactor
}
