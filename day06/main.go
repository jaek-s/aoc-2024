package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	log.Println("Puzzle 1 example output:", solveFirstPuzzle("./example.txt"))
	log.Println("Puzzle 1 real output:", solveFirstPuzzle("./input.txt"))
	log.Println("Puzzle 2 example output:", solveSecondPuzzle("./example.txt"))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type coord struct {
	x int
	y int
}

type turnPos struct {
	coord
	dir byte // Will always be 'N', 'E', 'S', or 'W'
}

type guardMap struct {
	width            int
	height           int
	obstacles        [][]bool
	startingPos      coord
	guardPos         coord
	guardDir         byte // Will always be 'N', 'E', 'S', or 'W'
	visitedLocations map[coord]bool
	turns            []turnPos
}

func newGuardMap(filepath string) guardMap {
	file, err := os.Open(filepath)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	guardMap := guardMap{
		guardDir:         'N',
		obstacles:        make([][]bool, 0),
		visitedLocations: make(map[coord]bool),
		turns:            make([]turnPos, 0),
	}
	var y int

	for scanner.Scan() {
		line := scanner.Text()
		if y == 0 {
			// The map input is always a square, so we can cheat here a little.
			mapWidth := len(line)
			guardMap.width = mapWidth
			guardMap.height = mapWidth

			guardMap.obstacles = make([][]bool, 0, mapWidth)
			for i := 0; i < mapWidth; i++ {
				guardMap.obstacles = append(guardMap.obstacles, make([]bool, mapWidth))
			}
		}

		for x, c := range line {
			if c == '#' {
				guardMap.obstacles[x][y] = true
			}

			if c == '^' {
				startingPos := coord{x: x, y: y}
				guardMap.guardPos = startingPos
				guardMap.startingPos = startingPos
				guardMap.visitedLocations[startingPos] = true
			}
		}

		y += 1
	}

	return guardMap
}

func (m *guardMap) moveGuard() {
	nextPos := m.getNextGuardPos()

	if nextPos.x >= 0 &&
		nextPos.x < m.width &&
		nextPos.y >= 0 &&
		nextPos.y < m.height &&
		m.obstacles[nextPos.x][nextPos.y] {

		m.rotateGuard()
		m.moveGuard()
		return
	}

	m.guardPos = nextPos
	m.visitedLocations[m.guardPos] = true
}

func (m *guardMap) rotateGuard() {
	m.turns = append(m.turns, turnPos{
		coord: coord{
			x: m.guardPos.x,
			y: m.guardPos.y,
		},
		dir: m.guardDir,
	})

	switch m.guardDir {
	case 'N':
		m.guardDir = 'E'
	case 'E':
		m.guardDir = 'S'
	case 'S':
		m.guardDir = 'W'
	case 'W':
		m.guardDir = 'N'
	}
}

func (m guardMap) getNextGuardPos() coord {
	newPos := m.guardPos

	if m.guardDir == 'N' {
		newPos.y -= 1
	}

	if m.guardDir == 'E' {
		newPos.x += 1
	}

	if m.guardDir == 'S' {
		newPos.y += 1
	}

	if m.guardDir == 'W' {
		newPos.x -= 1
	}

	return newPos
}

func (m guardMap) willNewObstacleLoopGuard(newObstacleLocation coord) bool {
	if newObstacleLocation == m.startingPos || m.visitedLocations[newObstacleLocation] {
		return false
	}

	m.obstacles[newObstacleLocation.x][newObstacleLocation.y] = true

	for m.guardPos.x < m.width && m.guardPos.y < m.height {
		m.moveGuard()
	}

	return false
}

func solveFirstPuzzle(filepath string) int {
	guardMap := newGuardMap(filepath)
	for guardMap.guardPos.x < guardMap.width && guardMap.guardPos.y < guardMap.height {
		guardMap.moveGuard()
	}

	// The final position will be off by one since the final off-the-map position is logged.ƒ
	return len(guardMap.visitedLocations) - 1
}

func solveSecondPuzzle(filepath string) int {
	guardMap := newGuardMap(filepath)

	for guardMap.guardPos.x < guardMap.width && guardMap.guardPos.y < guardMap.height {
		guardMap.moveGuard()
	}

	return 0
}
