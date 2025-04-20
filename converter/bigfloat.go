package converter

import (
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
)

// Function to convert big.Float to pgtype.Numeric for arithmetic operations with high accuracy and high precision
func BigFloatToPgNumeric(v *big.Float) *pgtype.Numeric {
	str := v.Text('f', -1)
	numeric := new(pgtype.Numeric)
	numeric.ScanScientific(str)

	return numeric
}

// Function to convert pgtype.Numeric to big.Float for arithmetic operations with high accuracy and high precision
func PgNumericToBigFloat(n pgtype.Numeric, precisionBits uint) *big.Float {
	// Check if the value is valid
	if !n.Valid {
		return nil
	}

	// Convert to big.Float
	intPart := new(big.Int).Set(n.Int)                                // Integer part
	floatVal := new(big.Float).SetInt(intPart).SetPrec(precisionBits) // Convert to big.Float
	exp := int64(n.Exp)                                               // Exponent

	// Calculate scale factor 10^exp
	ten := big.NewInt(10)
	scaleFactor := new(big.Float).SetPrec(precisionBits)

	if exp > 0 {
		// 10^exp
		scaleFactor.SetInt(new(big.Int).Exp(ten, big.NewInt(exp), nil))
		floatVal.Mul(floatVal, scaleFactor) // floatVal * (10^exp)
	} else if exp < 0 {
		// 10^(-exp)
		scaleFactor.SetInt(new(big.Int).Exp(ten, big.NewInt(-exp), nil))
		floatVal.Quo(floatVal, scaleFactor) // floatVal / (10^(-exp))
	}

	return floatVal
}
