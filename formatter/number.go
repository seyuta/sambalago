package formatter

import (
	"strconv"
	"strings"

	"github.com/seyuta/sambalago/converter"
)

// IndonesianNumber format number to indonesian number format (X.XXX) base on Pedoman Umum Ejaan Bahasa Indonesia (PUEBI) III.A.4
// https://badanbahasa.kemdikbud.go.id/lamanbahasa/sites/default/files/PUEBI.pdf
// Example: 1000000 to 1.000.000
func IndonesianNumber(x interface{}) string {
	r := converter.ToInt64(x)

	sign := ""
	if r < 0 {
		sign = "-"
		r = 0 - r
	}

	parts := []string{"", "", "", "", "", "", ""}
	j := len(parts) - 1

	for r > 999 {
		parts[j] = strconv.FormatInt(r%1000, 10)
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		r = r / 1000
		j--
	}
	parts[j] = strconv.Itoa(int(r))
	return sign + strings.Join(parts[j:], ".")
}
