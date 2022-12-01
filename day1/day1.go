package day1

import (
	"fmt"
	"strconv"

	t "github.com/alenius/aoctools"
)

type elfcarrier struct {
	index int
	// number of calories carried
	carries []string
	total   int
}

func Day1() {
	fmt.Println("day1")

	lines := t.ReadNewlineSeparatedFile("./day1/input.txt")

	elfs := []*elfcarrier{}
	for _, line := range lines {
		if len(elfs) == 0 {
			newElf := elfcarrier{index: len(elfs) + 1, carries: []string{}}
			elfs = append(elfs, &newElf)
		}

		if line == "" {
			newElf := elfcarrier{index: len(elfs) + 1, carries: []string{}}
			elfs = append(elfs, &newElf)
			continue
		}

		currentElf := elfs[len(elfs)-1]
		currentElf.carries = append(currentElf.carries, line)
		load, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		currentElf.total += load
	}

	leaders := [3]*elfcarrier{{}, {}, {}}
	for _, elf := range elfs {
		if elf.total > leaders[0].total {
			newLeaders := [3]*elfcarrier{elf}
			newLeaders[1] = leaders[0]
			newLeaders[2] = leaders[1]
			leaders = newLeaders
			continue
		}
		if elf.total > leaders[1].total {
			newLeaders := [3]*elfcarrier{leaders[0], elf}
			newLeaders[2] = leaders[1]
			leaders = newLeaders
			continue
		}
		if elf.total > leaders[1].total {
			newLeaders := [3]*elfcarrier{leaders[0], leaders[1], elf}
			leaders = newLeaders
			continue
		}
	}

	fmt.Printf("heaviest load: %v \n", leaders[0])
	fmt.Printf("top 3 total: %v \n", leaders[0].total+leaders[1].total+leaders[2].total)

	return
}
