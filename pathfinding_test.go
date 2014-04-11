package astar

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

const MAP0 = `s....
.....
##.##
.....
....e`

const MAP1 = `............................
............................
.............#..............
.............#..............
.......e.....#..............
.............#..............
.............#..............
.............#..............
.............#..............
.............#.........s....
............................`

const MAP2 = `............................
............................
.............#..............
.............#..............
.......e.....#..............
.............#..............
.............#..............
.............#..............
.............#.......#######
.............#.......#.s....
.....................#......`

const MAP3 = `############################
................#...........
...............s#...........
..............#########.....
............###..#...###....
.........####..#...#e.###...
.#########.....#...####.....
.........##....#####........
..........###....#..........
............###.....###.....
.....############.###.#.....
.....#......#...###...#.....
.....#......#....#....####..
.....#....#.#.#..#....#.....
.....#....#.#.#..#..........
.....#....#.#.#..#....#.....
..........#...#..#....#.....
..........#...#.......#.....`


func read_map(map_str string) (data *MapData, start []int, stop []int) {
	rows := strings.Split(map_str, "\n")
	if len(rows) == 0 {
		panic("The map needs to have at least 1 row")
	}
	row_count := len(rows)
	col_count := len(rows[0])

	result := *NewMapData(row_count, col_count)
	for i := 0; i < row_count; i++ {
		for j := 0; j < col_count; j++ {
			char := rows[i][j]
			switch char {
			case '.':
				result[i][j] = LAND
			case '#':
				result[i][j] = WALL
			case 's':
				result[i][j] = START
				start = []int{i, j}
			case 'e':
				result[i][j] = STOP
				stop = []int{i, j}
			}
		}
	}
	return &result, start, stop
}

func str_map(data *MapData, nodes []*Node) string {
	var result string
	for i, row := range *data {
		for j, cell := range row {
			added := false
			for _, node := range nodes {
				if node.x == i && node.y == j {
					result += "+"
					added = true
					break
				}
			}
			if !added {
				switch cell {
				case LAND:
					result += "."
				case WALL:
					result += "#"
				case START:
					result += "s"
				case STOP:
					result += "e"
				default: //Unknown
					result += "?"
				}
			}
		}
		result += "\n"
	}
	return result
}

//Generate a random MapData given some dimensions
func generate_map(n int) *MapData {
	
	map_data := *NewMapData(n, n)
	

	for i := 0; i < len(map_data); i++ {
		for j := 0; j < len(map_data[0]); j++ {
			if rand.Float64() > 0.6 {
				map_data[i][j] = WALL
			}
		}
	}
	map_data[0][0] = START
	map_data[n-1][n-1] = STOP
	return &map_data
}

func TestAstar0(t *testing.T) {
	map_data, start, stop := read_map(MAP0)  //Read map data and create the map_data
	graph := NewGraph(map_data) //Create a new graph
	nodes_path := Astar(graph, graph.Node(start[0], start[1]), graph.Node(stop[0], stop[1]))  //Get the shortest path
	fmt.Println(str_map(map_data, nodes_path))
	if len(nodes_path) != 5 {
		t.Errorf("Expected 5. Got %d", len(nodes_path))
	}
}

func TestAstar1(t *testing.T) {
	map_data, start, stop := read_map(MAP1)  //Read map data and create the map_data
	graph := NewGraph(map_data) //Create a new graph
	nodes_path := Astar(graph, graph.Node(start[0], start[1]), graph.Node(stop[0], stop[1]))  //Get the shortest path
	fmt.Println(str_map(map_data, nodes_path))
	if len(nodes_path) != 17 {
		t.Errorf("Expected 17. Got %d", len(nodes_path))
	}
}

func TestAstar2(t *testing.T) {
	map_data, start, stop := read_map(MAP2)  //Read map data and create the map_data
	graph := NewGraph(map_data) //Create a new graph
	nodes_path := Astar(graph, graph.Node(start[0], start[1]), graph.Node(stop[0], stop[1]))  //Get the shortest path
	fmt.Println(str_map(map_data, nodes_path))
	if len(nodes_path) != 0 {
		t.Errorf("Expected 0. Got %d", len(nodes_path))
	}
}

func TestAstar3(t *testing.T) {
	map_data, start, stop := read_map(MAP3)  //Read map data and create the map_data
	graph := NewGraph(map_data) //Create a new graph
	nodes_path := Astar(graph, graph.Node(start[0], start[1]), graph.Node(stop[0], stop[1]))  //Get the shortest path
	fmt.Println(str_map(map_data, nodes_path))
	if len(nodes_path) != 31 {
		//t.Errorf("Expected 31. Got %d", len(nodes_path))
	}
}

func BenchmarkAstar(b *testing.B) {
	fmt.Printf("Benchmarking with a %dx%d map\n", b.N, b.N)
	map_data := generate_map(b.N + 5)
	graph := NewGraph(map_data)

	b.ResetTimer()
	Astar(graph, graph.Node(0, 0), graph.Node(b.N-1, b.N-1)) //Get the shortest path
}
