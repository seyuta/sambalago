package formatter

// IndonesianRupiah format number to RpX.XXX base on Pedoman Umum Ejaan Bahasa Indonesia (PUEBI) II.H.5
// https://badanbahasa.kemdikbud.go.id/lamanbahasa/sites/default/files/PUEBI.pdf
func IndonesianRupiah(value interface{}) string {
	switch v := value.(type) {
	case int, int64, float64:
		return "Rp" + IndonesianNumber(v)
	default:
		return ""
	}
}
