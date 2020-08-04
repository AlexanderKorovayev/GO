package incapsulation_lib

import "errors"

type Date struct {
	year  int
	month int
	day   int
}

//так мы получаем лишь копию экземпляра Date
func (d Date) SetYear1(year int) {
	d.year = year
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) {
	d.month = month
}

func (d *Date) SetDay(day int) {
	d.day = day
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}
