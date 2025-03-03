package gjpd

import "time"

type NationalHoliday struct {
	Year  int
	Month int
	Day   int
	Date  string
	Name  string
}

func StringNextWeekday(date string) (t time.Time, e error) {
	d, err := time.Parse(time.DateOnly, date)
	if err != nil {
		e = err
		return
	}
	t = NextWeekday(d)
	return
}

func NextWeekday(d time.Time) (t time.Time) {

	startTime := d.AddDate(0, 0, 0)
	endTime := d.AddDate(1, 0, 0)

	for startTime.Compare(endTime) <= 0 {
		if IsHolyday(startTime) {
			startTime = startTime.AddDate(0, 0, 1)
		} else {
			t = startTime
			break
		}
	}

	return
}

func IsStringHolyday(date string) (b bool, e error) {
	d, err := time.Parse(time.DateOnly, date)
	if err != nil {
		e = err
		return
	}
	b = IsHolyday(d)
	return
}

func IsHolyday(d time.Time) (b bool) {
	weekday := int(d.Weekday())
	if weekday == 0 || weekday == 6 {
		b = true
		return
	}
	for _, h := range NationalHolidays() {
		if h.Year == d.Year() && h.Month == int(d.Month()) && h.Day == d.Day() {
			b = true
			break
		}
	}
	return
}

func IsStringNationalHolyday(date string) (b bool, e error) {
	d, err := time.Parse(time.DateOnly, date)
	if err != nil {
		e = err
		return
	}
	b = IsNationalHolyday(d)
	return
}

func IsNationalHolyday(d time.Time) (b bool) {
	for _, h := range NationalHolidays() {
		if h.Year == d.Year() && h.Month == int(d.Month()) && h.Day == d.Day() {
			b = true
		}
	}
	return
}

func NationalHolidays() []NationalHoliday {

	return []NationalHoliday{
		{
			Year:  2024,
			Month: 1,
			Day:   1,
			Date:  "2024-01-01",
			Name:  "元日",
		},
		{
			Year:  2024,
			Month: 1,
			Day:   8,
			Date:  "2024-01-08",
			Name:  "成人の日",
		},
		{
			Year:  2024,
			Month: 2,
			Day:   11,
			Date:  "2024-02-11",
			Name:  "建国記念の日",
		},
		{
			Year:  2024,
			Month: 2,
			Day:   12,
			Date:  "2024-02-12",
			Name:  "休日",
		},
		{
			Year:  2024,
			Month: 2,
			Day:   23,
			Date:  "2024-02-23",
			Name:  "天皇誕生日",
		},
		{
			Year:  2024,
			Month: 3,
			Day:   20,
			Date:  "2024-03-20",
			Name:  "春分の日",
		},
		{
			Year:  2024,
			Month: 4,
			Day:   29,
			Date:  "2024-04-29",
			Name:  "昭和の日",
		},
		{
			Year:  2024,
			Month: 5,
			Day:   3,
			Date:  "2024-05-03",
			Name:  "憲法記念日",
		},
		{
			Year:  2024,
			Month: 5,
			Day:   4,
			Date:  "2024-05-04",
			Name:  "みどりの日",
		},
		{
			Year:  2024,
			Month: 5,
			Day:   5,
			Date:  "2024-05-05",
			Name:  "こどもの日",
		},
		{
			Year:  2024,
			Month: 5,
			Day:   6,
			Date:  "2024-05-06",
			Name:  "休日",
		},
		{
			Year:  2024,
			Month: 7,
			Day:   15,
			Date:  "2024-07-15",
			Name:  "海の日",
		},
		{
			Year:  2024,
			Month: 8,
			Day:   11,
			Date:  "2024-08-11",
			Name:  "山の日",
		},
		{
			Year:  2024,
			Month: 8,
			Day:   12,
			Date:  "2024-08-12",
			Name:  "休日",
		},
		{
			Year:  2024,
			Month: 9,
			Day:   16,
			Date:  "2024-09-16",
			Name:  "敬老の日",
		},
		{
			Year:  2024,
			Month: 9,
			Day:   22,
			Date:  "2024-09-22",
			Name:  "秋分の日",
		},
		{
			Year:  2024,
			Month: 9,
			Day:   23,
			Date:  "2024-09-23",
			Name:  "休日",
		},
		{
			Year:  2024,
			Month: 10,
			Day:   14,
			Date:  "2024-10-14",
			Name:  "スポーツの日",
		},
		{
			Year:  2024,
			Month: 11,
			Day:   3,
			Date:  "2024-11-03",
			Name:  "文化の日",
		},
		{
			Year:  2024,
			Month: 11,
			Day:   4,
			Date:  "2024-11-04",
			Name:  "休日",
		},
		{
			Year:  2024,
			Month: 11,
			Day:   23,
			Date:  "2024-11-23",
			Name:  "勤労感謝の日",
		},
		{
			Year:  2025,
			Month: 1,
			Day:   1,
			Date:  "2025-01-01",
			Name:  "元日",
		},
		{
			Year:  2025,
			Month: 1,
			Day:   13,
			Date:  "2025-01-13",
			Name:  "成人の日",
		},
		{
			Year:  2025,
			Month: 2,
			Day:   11,
			Date:  "2025-02-11",
			Name:  "建国記念の日",
		},
		{
			Year:  2025,
			Month: 2,
			Day:   23,
			Date:  "2025-02-23",
			Name:  "天皇誕生日",
		},
		{
			Year:  2025,
			Month: 2,
			Day:   24,
			Date:  "2025-02-24",
			Name:  "休日",
		},
		{
			Year:  2025,
			Month: 3,
			Day:   20,
			Date:  "2025-03-20",
			Name:  "春分の日",
		},
		{
			Year:  2025,
			Month: 4,
			Day:   29,
			Date:  "2025-04-29",
			Name:  "昭和の日",
		},
		{
			Year:  2025,
			Month: 5,
			Day:   3,
			Date:  "2025-05-03",
			Name:  "憲法記念日",
		},
		{
			Year:  2025,
			Month: 5,
			Day:   4,
			Date:  "2025-05-04",
			Name:  "みどりの日",
		},
		{
			Year:  2025,
			Month: 5,
			Day:   5,
			Date:  "2025-05-05",
			Name:  "こどもの日",
		},
		{
			Year:  2025,
			Month: 5,
			Day:   6,
			Date:  "2025-05-06",
			Name:  "休日",
		},
		{
			Year:  2025,
			Month: 7,
			Day:   21,
			Date:  "2025-07-21",
			Name:  "海の日",
		},
		{
			Year:  2025,
			Month: 8,
			Day:   11,
			Date:  "2025-08-11",
			Name:  "山の日",
		},
		{
			Year:  2025,
			Month: 9,
			Day:   15,
			Date:  "2025-09-15",
			Name:  "敬老の日",
		},
		{
			Year:  2025,
			Month: 9,
			Day:   23,
			Date:  "2025-09-23",
			Name:  "秋分の日",
		},
		{
			Year:  2025,
			Month: 10,
			Day:   13,
			Date:  "2025-10-13",
			Name:  "スポーツの日",
		},
		{
			Year:  2025,
			Month: 11,
			Day:   3,
			Date:  "2025-11-03",
			Name:  "文化の日",
		},
		{
			Year:  2025,
			Month: 11,
			Day:   23,
			Date:  "2025-11-23",
			Name:  "勤労感謝の日",
		},
		{
			Year:  2025,
			Month: 11,
			Day:   24,
			Date:  "2025-11-24",
			Name:  "休日",
		},

		{
			Year:  2026,
			Month: 1,
			Day:   1,
			Date:  "2026-01-01",
			Name:  "元日",
		},
		{
			Year:  2026,
			Month: 1,
			Day:   12,
			Date:  "2026-01-12",
			Name:  "成人の日",
		},
		{
			Year:  2026,
			Month: 2,
			Day:   11,
			Date:  "2026-02-11",
			Name:  "建国記念の日",
		},
		{
			Year:  2026,
			Month: 2,
			Day:   23,
			Date:  "2026-02-23",
			Name:  "天皇誕生日",
		},
		{
			Year:  2026,
			Month: 3,
			Day:   20,
			Date:  "2026-03-20",
			Name:  "春分の日",
		},
		{
			Year:  2026,
			Month: 4,
			Day:   29,
			Date:  "2026-04-29",
			Name:  "昭和の日",
		},
		{
			Year:  2026,
			Month: 5,
			Day:   3,
			Date:  "2026-05-03",
			Name:  "憲法記念日",
		},
		{
			Year:  2026,
			Month: 5,
			Day:   4,
			Date:  "2026-05-04",
			Name:  "みどりの日",
		},
		{
			Year:  2026,
			Month: 5,
			Day:   5,
			Date:  "2026-05-05",
			Name:  "こどもの日",
		},
		{
			Year:  2026,
			Month: 5,
			Day:   6,
			Date:  "2026-05-06",
			Name:  "休日",
		},
		{
			Year:  2026,
			Month: 7,
			Day:   20,
			Date:  "2026-07-20",
			Name:  "海の日",
		},
		{
			Year:  2026,
			Month: 8,
			Day:   11,
			Date:  "2026-08-11",
			Name:  "山の日",
		},
		{
			Year:  2026,
			Month: 9,
			Day:   21,
			Date:  "2026-09-21",
			Name:  "敬老の日",
		},
		{
			Year:  2026,
			Month: 9,
			Day:   22,
			Date:  "2026-09-22",
			Name:  "休日",
		},
		{
			Year:  2026,
			Month: 9,
			Day:   23,
			Date:  "2026-09-23",
			Name:  "秋分の日",
		},
		{
			Year:  2026,
			Month: 10,
			Day:   12,
			Date:  "2026-10-12",
			Name:  "スポーツの日",
		},
		{
			Year:  2026,
			Month: 11,
			Day:   3,
			Date:  "2026-11-03",
			Name:  "文化の日",
		},
		{
			Year:  2026,
			Month: 11,
			Day:   23,
			Date:  "2026-11-23",
			Name:  "勤労感謝の日",
		},
	}
}
