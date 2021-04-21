package main

import "fmt"

type Byte int8
type Word int16

type CPU struct {
	PC Word // program counter
	SP Byte // stack pointer

	A, X, Y Byte // registers

	C, Z, I, D, B, V, N Byte // status flag
}

func (cpu *CPU) InitFlags() {
	cpu.C = 1
	cpu.Z = 1
	cpu.I = 1
	cpu.D = 1
	cpu.B = 1
	cpu.V = 1
	cpu.N = 1
}

func (cpu *CPU) Reset() {
	cpu.PC = 4092 // 0xFFFC
	cpu.SP = 4095 // 0x0100

	cpu.A, cpu.X, cpu.Y = 0, 0, 0 // registers

	cpu.C, cpu.Z, cpu.I, cpu.D = 0, 0, 0, 0
	cpu.B, cpu.V, cpu.N = 0, 0, 0

}

func main() {
	cpu := CPU{}
	cpu.InitFlags()
	cpu.Reset()

	fmt.Println(cpu.Z)

}
