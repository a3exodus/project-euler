package helpers

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func initG() map[string][]Edge {
	file, err := os.Open("graphs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	G := map[string][]Edge{}

	scanner := bufio.NewScanner(file)
	v := 1
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, `,`)
		for w, num := range numbers {
			val, _ := strconv.Atoi(num)
			if val != 0 {
				G[strconv.Itoa(v)] = append(
					G[strconv.Itoa(v)],
					Edge{strconv.Itoa(w + 1), val},
				)
			}
		}
		v++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return G
}

func TestDijkstrasAlgorithm(t *testing.T) {
	G := initG()
	dist := DijkstrasAlgorithm(G, "1")
	assert.Equal(t, 4, dist["9"])
}
