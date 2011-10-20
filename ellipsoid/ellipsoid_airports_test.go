package ellipsoid

import "fmt"
import "math"

var pass, fail int

func delta_within(v float64, t float64, d float64) {
	diff := math.Fabs(v - t)
	if diff < d {
		fmt.Printf("Test OK\n")
		pass++
	} else {
		fmt.Printf("Test FAIL\n")
		fail++
		fmt.Printf("v - t = %f - %f = %f\n", v, t, diff)
	}
	fmt.Printf("---\n")
}

func main() {
	epsilon := 50.0

	e := ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Kilometer, ellipsoid.Longitude_is_symmetric, ellipsoid.Bearing_is_symmetric)

	fmt.Printf("Going from Istanbul to Delhi = 4550\n")
	r, _ := e.To(41.1, 29, 28.67, 77.21)
	delta_within(r, 4550, epsilon)

	fmt.Printf("Going from Tokyo to New_York = 10846\n")
	r, _ = e.To(35.67, 139.77, 40.67, -73.94)
	delta_within(r, 10846, epsilon)

	fmt.Printf("Going from Lima to Cairo = 12423\n")
	r, _ = e.To(-12.07, -77.05, 30.06, 31.25)
	delta_within(r, 12423, epsilon)

	fmt.Printf("Going from Lahore to Rio_de_Janeiro = 13845\n")
	r, _ = e.To(31.56, 74.35, -22.91, -43.2)
	delta_within(r, 13845, epsilon)

	fmt.Printf("Going from Santiago to Kaifeng = 19529\n")
	r, _ = e.To(19.48, -70.69, 34.85, 114.35)
	delta_within(r, 19529, epsilon)

	fmt.Printf("Going from Calcutta to Toronto = 12540\n")
	r, _ = e.To(22.57, 88.36, 43.65, -79.38)
	delta_within(r, 12540, epsilon)

	fmt.Printf("Going from Rangoon to Sydney = 8104\n")
	r, _ = e.To(16.79, 96.15, -33.87, 151.21)
	delta_within(r, 8104, epsilon)

	fmt.Printf("Going from Madras to Riyadh = 3739\n")
	r, _ = e.To(13.09, 80.27, 24.65, 46.77)
	delta_within(r, 3739, epsilon)

	fmt.Printf("Going from Chongqing to Chengdu = 268\n")
	r, _ = e.To(29.57, 106.58, 30.67, 104.07)
	delta_within(r, 268, epsilon)

	fmt.Printf("Going from Tianjin to Melbourne = 9015\n")
	r, _ = e.To(39.13, 117.2, -37.81, 144.96)
	delta_within(r, 9015, epsilon)

	fmt.Printf("Going from Pusan to Abidjan = 13362\n")
	r, _ = e.To(35.11, 129.03, 5.33, -4.03)
	delta_within(r, 13362, epsilon)

	fmt.Printf("Going from Yokohama to Ibadan = 13373\n")
	r, _ = e.To(35.47, 139.62, 7.38, 3.93)
	delta_within(r, 13373, epsilon)

	fmt.Printf("Going from Singapore to Ankara = 8303\n")
	r, _ = e.To(1.3, 103.85, 39.93, 32.85)
	delta_within(r, 8303, epsilon)

	fmt.Printf("Going from Berlin to Montreal = 6001\n")
	r, _ = e.To(52.52, 13.38, 45.52, -73.57)
	delta_within(r, 6001, epsilon)

	fmt.Printf("Going from Pyongyang to Lanzhou = 1959\n")
	r, _ = e.To(39.02, 125.75, 36.05, 103.68)
	delta_within(r, 1959, epsilon)

	fmt.Printf("Going from Guangzhou to Casablanca = 11132\n")
	r, _ = e.To(23.12, 113.25, 33.6, -7.62)
	delta_within(r, 11132, epsilon)

	fmt.Printf("Going from Durban to Madrid = 8593\n")
	r, _ = e.To(-29.87, 30.99, 40.42, -3.71)
	delta_within(r, 8593, epsilon)

	fmt.Printf("Going from Nanjing to Kabul = 4571\n")
	r, _ = e.To(32.05, 118.78, 34.53, 69.17)
	delta_within(r, 4571, epsilon)

	fmt.Printf("Going from Pune to Surat = 312\n")
	r, _ = e.To(18.53, 73.84, 21.2, 72.82)
	delta_within(r, 312, epsilon)

	fmt.Printf("Going from Jiddah to Chicago = 11102\n")
	r, _ = e.To(21.5, 39.17, 41.84, -87.68)
	delta_within(r, 11102, epsilon)

	fmt.Printf("Going from Kanpur to Luanda = 8228\n")
	r, _ = e.To(26.47, 80.33, -8.82, 13.24)
	delta_within(r, 8228, epsilon)

	fmt.Printf("Going from Taiyuan to Salvador = 16033\n")
	r, _ = e.To(37.87, 112.55, -12.97, -38.5)
	delta_within(r, 16033, epsilon)

	fmt.Printf("Going from Taegu to Rome = 9202\n")
	r, _ = e.To(35.87, 128.6, 41.89, 12.5)
	delta_within(r, 9202, epsilon)

	fmt.Printf("Going from Changchun to Kiev = 6701\n")
	r, _ = e.To(43.87, 125.35, 50.43, 30.52)
	delta_within(r, 6701, epsilon)

	fmt.Printf("Going from Faisalabad to Izmir = 4215\n")
	r, _ = e.To(31.41, 73.11, 38.43, 27.15)
	delta_within(r, 4215, epsilon)

	fmt.Printf("Summary: pass=%v fail=%v\n", pass, fail)

}
