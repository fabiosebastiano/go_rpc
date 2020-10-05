package sample

import (
	"math/rand"
	"time"

	"github.com/fabiosebastiano/go_grpc/pb"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func RandomLaptopName(brand string) string {
	switch brand {
	case "APPLE":
		return randomStringFromSet("Macbook Air", "Macbook Pro")
	case "DELL":
		return randomStringFromSet("Latitude", "Vostro", "XPS")
	default:
		return randomStringFromSet("Thinkpad X1", "Thinkpad P1", "Thinkpad P53")
	}
}

func RandomBool() bool {
	return rand.Intn(2) == 1
}

func RandomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func RandomCPUBrand() string {
	return randomStringFromSet("INTEL", "AMD")
}

func RandomCPUName(brand string) string {
	if brand == "INTEL" {
		return randomStringFromSet(
			"Xeon E",
			"Core i9",
			"Core i7",
			"Core i11",
			"Core i14",
		)
	}
	return randomStringFromSet(
		"Ryzen 1",
		"Ryzen 2",
		"Ryzen 3",
		"Ryzen 4",
	)
}

func RandomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func RandomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"RTX 1",
			"RTX 2",
			"RTX 3",
			"RTX 4",
		)
	}
	return randomStringFromSet(
		"RX 1",
		"RX 2",
		"RX 3",
		"RX 4",
	)
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomInt(min int, max int) int {

	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
