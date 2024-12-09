package moneylib

import (
	"encoding/json"
	"errors"
	"log"
)

var currencyJSON string = `[
  {"code": "AED", "sym": "د.إ", "number": 784, "nod": 2},
  {"code": "AFN", "sym": "؋", "number": 971, "nod": 2},
  {"code": "ALL", "sym": "L", "number": 8, "nod": 2},
  {"code": "AMD", "sym": "դր.", "number": 51, "nod": 2},
  {"code": "ANG", "sym": "ƒ", "number": 532, "nod": 2},
  {"code": "AOA", "sym": "Kz", "number": 973, "nod": 2},
  {"code": "ARS", "sym": "$", "number": 32, "nod": 2},
  {"code": "AUD", "sym": "$", "number": 36, "nod": 2},
  {"code": "AWG", "sym": "ƒ", "number": 533, "nod": 2},
  {"code": "AZN", "sym": "₼", "number": 944, "nod": 2},
  {"code": "BAM", "sym": "KM", "number": 977, "nod": 2},
  {"code": "BBD", "sym": "$", "number": 52, "nod": 2},
  {"code": "BDT", "sym": "৳", "number": 50, "nod": 2},
  {"code": "BGN", "sym": "лв", "number": 975, "nod": 2},
  {"code": "BHD", "sym": "ب.د", "number": 48, "nod": 3},
  {"code": "BIF", "sym": "Fr", "number": 108, "nod": 0},
  {"code": "BMD", "sym": "$", "number": 60, "nod": 2},
  {"code": "BND", "sym": "$", "number": 96, "nod": 2},
  {"code": "BOB", "sym": "Bs.", "number": 68, "nod": 2},
  {"code": "BOV", "sym": "", "number": 984, "nod": 2},
  {"code": "BRL", "sym": "R$", "number": 986, "nod": 2},
  {"code": "BSD", "sym": "$", "number": 44, "nod": 2},
  {"code": "BTN", "sym": "Nu.", "number": 64, "nod": 2},
  {"code": "BWP", "sym": "P", "number": 72, "nod": 2},
  {"code": "BYN", "sym": "Br", "number": 933, "nod": 2},
  {"code": "BZD", "sym": "$", "number": 84, "nod": 2},
  {"code": "CAD", "sym": "$", "number": 124, "nod": 2},
  {"code": "CDF", "sym": "Fr", "number": 976, "nod": 2},
  {"code": "CHE", "sym": "", "number": 947, "nod": 2},
  {"code": "CHF", "sym": "Fr", "number": 756, "nod": 2},
  {"code": "CHW", "sym": "", "number": 948, "nod": 2},
  {"code": "CLF", "sym": "UF", "number": 990, "nod": 4},
  {"code": "CLP", "sym": "$", "number": 152, "nod": 0},
  {"code": "CNY", "sym": "¥", "number": 156, "nod": 2},
  {"code": "COP", "sym": "$", "number": 170, "nod": 2},
  {"code": "COU", "sym": "", "number": 970, "nod": 2},
  {"code": "CRC", "sym": "₡", "number": 188, "nod": 2},
  {"code": "CUC", "sym": "$", "number": 931, "nod": 2},
  {"code": "CUP", "sym": "$", "number": 192, "nod": 2},
  {"code": "CVE", "sym": "$", "number": 132, "nod": 2},
  {"code": "CZK", "sym": "Kč", "number": 203, "nod": 2},
  {"code": "DJF", "sym": "Fdj", "number": 262, "nod": 0},
  {"code": "DKK", "sym": "kr", "number": 208, "nod": 2},
  {"code": "DOP", "sym": "$", "number": 214, "nod": 2},
  {"code": "DZD", "sym": "د.ج", "number": 12, "nod": 2},
  {"code": "EGP", "sym": "KR", "number": 818, "nod": 2},
  {"code": "ERN", "sym": "", "number": 232, "nod": 2},
  {"code": "ETB", "sym": "", "number": 230, "nod": 2},
  {"code": "EUR", "sym": "€", "number": 978, "nod": 2},
  {"code": "FJD", "sym": "$", "number": 242, "nod": 2},
  {"code": "FKP", "sym": "£", "number": 238, "nod": 2},
  {"code": "GBP", "sym": "£", "number": 826, "nod": 2},
  {"code": "GEL", "sym": "ლ", "number": 981, "nod": 2},
  {"code": "GHS", "sym": "₵", "number": 936, "nod": 2},
  {"code": "GIP", "sym": "£", "number": 292, "nod": 2},
  {"code": "GMD", "sym": "D", "number": 270, "nod": 2},
  {"code": "GNF", "sym": "Fr", "number": 324, "nod": 0},
  {"code": "GTQ", "sym": "Q", "number": 320, "nod": 2},
  {"code": "GYD", "sym": "$", "number": 328, "nod": 2},
  {"code": "HKD", "sym": "$", "number": 344, "nod": 2},
  {"code": "HNL", "sym": "L", "number": 340, "nod": 2},
  {"code": "HRK", "sym": "kn", "number": 191, "nod": 2},
  {"code": "HTG", "sym": "G", "number": 332, "nod": 2},
  {"code": "HUF", "sym": "Ft", "number": 348, "nod": 2},
  {"code": "IDR", "sym": "Rp", "number": 360, "nod": 2},
  {"code": "ILS", "sym": "₪", "number": 376, "nod": 2},
  {"code": "INR", "sym": "₹", "number": 356, "nod": 2},
  {"code": "IQD", "sym": "ع.د", "number": 368, "nod": 3},
  {"code": "IRR", "sym": "﷼", "number": 364, "nod": 2},
  {"code": "ISK", "sym": "kr", "number": 352, "nod": 0},
  {"code": "JMD", "sym": "$", "number": 388, "nod": 2},
  {"code": "JOD", "sym": "د.ا", "number": 400, "nod": 3},
  {"code": "JPY", "sym": "¥", "number": 392, "nod": 0},
  {"code": "KES", "sym": "KSh", "number": 404, "nod": 2},
  {"code": "KGS", "sym": "som", "number": 417, "nod": 2},
  {"code": "KHR", "sym": "៛", "number": 116, "nod": 2},
  {"code": "KMF", "sym": "Fr", "number": 174, "nod": 0},
  {"code": "KPW", "sym": "₩", "number": 408, "nod": 2},
  {"code": "KRW", "sym": "₩", "number": 410, "nod": 0},
  {"code": "KWD", "sym": "د.ك", "number": 414, "nod": 3},
  {"code": "KYD", "sym": "$", "number": 136, "nod": 2},
  {"code": "KZT", "sym": "〒", "number": 398, "nod": 2},
  {"code": "LAK", "sym": "₭", "number": 418, "nod": 2},
  {"code": "LBP", "sym": "ل.ل", "number": 422, "nod": 2},
  {"code": "LKR", "sym": "₨", "number": 144, "nod": 2},
  {"code": "LRD", "sym": "$", "number": 430, "nod": 2},
  {"code": "LSL", "sym": "L", "number": 426, "nod": 2},
  {"code": "LYD", "sym": "", "number": 434, "nod": 3},
  {"code": "MAD", "sym": "د.م.", "number": 504, "nod": 2},
  {"code": "MDL", "sym": "L", "number": 498, "nod": 2},
  {"code": "MGA", "sym": "Ar", "number": 969, "nod": 2},
  {"code": "MKD", "sym": "ден", "number": 807, "nod": 2},
  {"code": "MMK", "sym": "K", "number": 104, "nod": 2},
  {"code": "MNT", "sym": "₮", "number": 496, "nod": 2},
  {"code": "MOP", "sym": "P", "number": 446, "nod": 2},
  {"code": "MRO", "sym": "UM", "number": 478, "nod": 2},
  {"code": "MUR", "sym": "₨", "number": 480, "nod": 2},
  {"code": "MVR", "sym": "MVR", "number": 462, "nod": 2},
  {"code": "MWK", "sym": "MK", "number": 454, "nod": 2},
  {"code": "MXN", "sym": "$", "number": 484, "nod": 2},
  {"code": "MXV", "sym": "", "number": 979, "nod": 2},
  {"code": "MYR", "sym": "RM", "number": 458, "nod": 2},
  {"code": "MZN", "sym": "MTn", "number": 943, "nod": 2},
  {"code": "NAD", "sym": "$", "number": 516, "nod": 2},
  {"code": "NGN", "sym": "₦", "number": 566, "nod": 2},
  {"code": "NIO", "sym": "C$", "number": 558, "nod": 2},
  {"code": "NOK", "sym": "kr", "number": 578, "nod": 2},
  {"code": "NPR", "sym": "₨", "number": 524, "nod": 2},
  {"code": "NZD", "sym": "$", "number": 554, "nod": 2},
  {"code": "OMR", "sym": "ر.ع.", "number": 512, "nod": 3},
  {"code": "PAB", "sym": "B/.", "number": 590, "nod": 2},
  {"code": "PEN", "sym": "S/.", "number": 604, "nod": 2},
  {"code": "PGK", "sym": "K", "number": 598, "nod": 2},
  {"code": "PHP", "sym": "₱", "number": 608, "nod": 2},
  {"code": "PKR", "sym": "₨", "number": 586, "nod": 2},
  {"code": "PLN", "sym": "zł", "number": 985, "nod": 2},
  {"code": "PYG", "sym": "₲", "number": 600, "nod": 0},
  {"code": "QAR", "sym": "ر.ق", "number": 634, "nod": 2},
  {"code": "RON", "sym": "Lei", "number": 946, "nod": 2},
  {"code": "RSD", "sym": "РСД", "number": 941, "nod": 2},
  {"code": "RUB", "sym": "₽", "number": 643, "nod": 2},
  {"code": "RWF", "sym": "FRw", "number": 646, "nod": 0},
  {"code": "SAR", "sym": "ر.س", "number": 682, "nod": 2},
  {"code": "SBD", "sym": "$", "number": 90, "nod": 2},
  {"code": "SCR", "sym": "₨", "number": 690, "nod": 2},
  {"code": "SDG", "sym": "£", "number": 938, "nod": 2},
  {"code": "SEK", "sym": "kr", "number": 752, "nod": 2},
  {"code": "SGD", "sym": "$", "number": 702, "nod": 2},
  {"code": "SHP", "sym": "£", "number": 654, "nod": 2},
  {"code": "SLL", "sym": "Le", "number": 694, "nod": 2},
  {"code": "SOS", "sym": "Sh", "number": 706, "nod": 2},
  {"code": "SRD", "sym": "$", "number": 968, "nod": 2},
  {"code": "SSP", "sym": "£", "number": 728, "nod": 2},
  {"code": "STD", "sym": "Db", "number": 678, "nod": 2},
  {"code": "SVC", "sym": "₡", "number": 222, "nod": 2},
  {"code": "SYP", "sym": "£S", "number": 760, "nod": 2},
  {"code": "SZL", "sym": "L", "number": 748, "nod": 2},
  {"code": "THB", "sym": "฿", "number": 764, "nod": 2},
  {"code": "TJS", "sym": "SM", "number": 972, "nod": 2},
  {"code": "TMT", "sym": "T", "number": 934, "nod": 2},
  {"code": "TND", "sym": "د.ت", "number": 788, "nod": 3},
  {"code": "TOP", "sym": "T$", "number": 776, "nod": 2},
  {"code": "TRY", "sym": "₺", "number": 949, "nod": 2},
  {"code": "TTD", "sym": "$", "number": 780, "nod": 2},
  {"code": "TWD", "sym": "$", "number": 901, "nod": 2},
  {"code": "TZS", "sym": "Sh", "number": 834, "nod": 2},
  {"code": "UAH", "sym": "₴", "number": 980, "nod": 2},
  {"code": "UGX", "sym": "USh", "number": 800, "nod": 0},
  {"code": "USD", "sym": "$", "number": 840, "nod": 2},
  {"code": "USN", "sym": "", "number": 997, "nod": 2},
  {"code": "UYI", "sym": "", "number": 940, "nod": 0},
  {"code": "UYU", "sym": "$", "number": 858, "nod": 2},
  {"code": "UZS", "sym": "лв", "number": 860, "nod": 2},
  {"code": "VEF", "sym": "Bs F", "number": 937, "nod": 2},
  {"code": "VND", "sym": "₫", "number": 704, "nod": 0},
  {"code": "VUV", "sym": "Vt", "number": 548, "nod": 0},
  {"code": "WST", "sym": "T", "number": 882, "nod": 2},
  {"code": "XAF", "sym": "Fr", "number": 950, "nod": 0},
  {"code": "XAG", "sym": "oz t", "number": 961, "nod": null},
  {"code": "XAU", "sym": "oz t", "number": 959, "nod": null},
  {"code": "XBA", "sym": "", "number": 955, "nod": null},
  {"code": "XBB", "sym": "", "number": 956, "nod": null},
  {"code": "XBC", "sym": "", "number": 957, "nod": null},
  {"code": "XBD", "sym": "", "number": 958, "nod": null},
  {"code": "XCD", "sym": "$", "number": 951, "nod": 2},
  {"code": "XDR", "sym": "SDR", "number": 960, "nod": null},
  {"code": "XOF", "sym": "Fr", "number": 952, "nod": 0},
  {"code": "XPD", "sym": "", "number": 964, "nod": null},
  {"code": "XPF", "sym": "Fr", "number": 953, "nod": 0},
  {"code": "XPT", "sym": "", "number": 962, "nod": null},
  {"code": "XSU", "sym": "", "number": 994, "nod": null},
  {"code": "XTS", "sym": "", "number": 963, "nod": null},
  {"code": "XUA", "sym": "", "number": 965, "nod": null},
  {"code": "XXX", "sym": "", "number": 999, "nod": null},
  {"code": "YER", "sym": "﷼", "number": 886, "nod": 2},
  {"code": "ZAR", "sym": "R", "number": 710, "nod": 2},
  {"code": "ZMW", "sym": "ZK", "number": 967, "nod": 2},
  {"code": "ZWL", "sym": "$", "number": 932, "nod": 2}
]
`

var currMap = make(map[string]Currency)

func init() {
	currencies := make([]currTmp, 0)
	err := json.Unmarshal([]byte(currencyJSON), &currencies)
	if err != nil {
		log.Fatalf("Error unmarshaling currency data")
	}
	for _, v := range currencies {
		currMap[v.Code] = newCurrency(&v)
	}
}

// GetCurrency returns the currency object for the code.
func GetCurrency(code string) (Currency, error) {
	curr := currMap[code]
	if curr.code != code {
		return curr, errors.New("Currency code not available: " + code)
	}
	return currMap[code], nil
}

// ValidCurrency checks if the currency code is correct
func ValidCurrency(code string) bool {
	_, err := GetCurrency(code)
	return err == nil
}