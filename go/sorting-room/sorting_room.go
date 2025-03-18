package sorting

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(NumberBox.Number(nb)))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	fancy, ok := fnb.(FancyNumber)
	if !ok {
		return 0
	}
	value, err := strconv.Atoi(fancy.Value())
	if err != nil {
		return 0
	}
	return value
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	switch this := i.(type) {
	case string:
		return "Return to sender"
	case int:
		return DescribeNumber(float64(this))
	case float64:
		return DescribeNumber(this)
	case NumberBox:
		return DescribeNumberBox(this)
	case FancyNumber:
		return DescribeFancyNumberBox(this)
	case FancyNumberBox:
		return DescribeFancyNumberBox(this)
	default:
		return ""
	}
}
