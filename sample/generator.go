package sample

import (
	"github.com/fabiosebastiano/go_grpc/pb"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
)

// NewKeyboard ritorna
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  RandomKeyboardLayout(),
		Backlit: RandomBool(),
	}

	return keyboard
}

// NewCPU
func NewCPU() *pb.CPU {

	brand := RandomCPUBrand()
	name := RandomCPUName(brand)
	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5)

	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}

	return cpu
}

// NewGPU
func NewGPU() *pb.GPU {
	brand := RandomGPUBrand()
	name := RandomGPUName(brand)

	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)

	memory := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}

	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}
	return gpu
}

// NewRAM
func NewRAM() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(randomInt(6, 64)),
		Unit:  pb.Memory_GIGABYTE,
	}
	return ram
}

// NewSSD
func NewSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	return ssd
}

// NewHDD
func NewHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}
	return hdd
}

func NewScreen() *pb.Screen {
	height := randomInt(1080, 4320)
	width := height * 16 / 9

	screen := &pb.Screen{
		SizeInch: randomFloat32(13, 64),
		Resolution: &pb.Screen_Resolution{
			Height: uint32(height),
			Width:  uint32(width),
		},
		Panel:      RandomScreenPanel(),
		Multitouch: RandomBool(),
	}

	return screen
}

func NewLaptop() *pb.Laptop {
	brand := randomStringFromSet("APPLE", "DELL", "LENOVO")
	laptop := &pb.Laptop{
		Id:   uuid.New().String(),
		Brand: brand,
		Name: RandomLaptopName(brand),
		Cpu: NewCPU(),
		Ram: NewRAM(),
		Gpus: []*pb.GPU{NewGPU()},
		Storages: []*pb.Storage{NewSSD(), NewHDD()},
		Screen: NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd: randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2001,2020)) ,
		UpdatedAt: ptypes.TimestampNow(),
	}

	return laptop
}
