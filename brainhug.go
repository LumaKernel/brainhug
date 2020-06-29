package brainhug

import (
	"log"
	"errors"

	"github.com/LumaKernel/brainhug/internal/cycle"
)

var Logging = false

const (
	// brainfuck keywords
	kw_next = '>'
	kw_prev = '<'
	kw_read = ','
	kw_wirte = '.'
	kw_count_up = '+'
	kw_count_down = '-'
	kw_loop_start = '['
	kw_loop_end = ']'

	err_loop_not_match = "Number of loop ends should equal to that of loop starts."
)

func Proceed(programStr string, stdin []byte) ([]byte, error) {
	program := []rune(programStr)

	ptr := int64(0)
	in_ptr := 0

	mem := new(cycle.Cycle)
	mem.AppendFront(0)
	buf := []byte{}


	pc := 0 // program counter

	for pc < len(program) {
		if Logging {
			log.Printf("%-5d %4c %6d %6d %v\n", pc, program[pc], ptr, mem.Get(ptr), buf)
		}
		switch program[pc] {
		case kw_next:
			ptr++
			if ptr >= mem.LenPos() {
				mem.AppendFront(0)
			}
		case kw_prev:
			ptr--
			if -ptr > mem.LenNeg() {
				mem.AppendBack(0)
			}
		case kw_read:
			r := byte(0)
			if in_ptr < len(stdin) {
				r = stdin[in_ptr]
				in_ptr++
			}
			mem.Set(ptr, r)
		case kw_wirte:
			buf = append(buf, mem.Get(ptr))
		case kw_count_up:
			mem.Set(ptr, mem.Get(ptr) + 1)
		case kw_count_down:
			mem.Set(ptr, mem.Get(ptr) - 1)
		case kw_loop_start:
			if mem.Get(ptr) == 0 {
				jumping := 1
				pc++
				forward:
				for pc < len(program) {
					switch program[pc] {
					case kw_loop_start:
						jumping++
					case kw_loop_end:
						jumping--
						if jumping == 0 {
							break forward
						}
					}
					pc++
				}
				if jumping > 0 {
					return nil, errors.New(err_loop_not_match)
				}
				pc++
				continue
			}
		case kw_loop_end:
			jumping := 1
			pc--
			back:
			for pc >= 0 {
				switch program[pc] {
				case kw_loop_start:
					jumping--
					if jumping == 0 {
						break back
					}
				case kw_loop_end:
					jumping++
				}
				pc--
			}
			if jumping > 0 {
				return nil, errors.New(err_loop_not_match)
			}
			continue
		}
		pc++
	}
	return buf, nil
}

