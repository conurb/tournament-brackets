package main

import (
	"fmt"
	"strings"
)

var tournament1 = []string{"BIG DIGS", "WIG ZIG", "BIG FIGURES", "TIGGER WIG", "FIGURES TIGGER", "TIGGER"}

/* Tournament 1
_BIG__
      \_BIG_____
_DIGS_/         \
                 \_FIGURES_
                 /         \
       _FIGURES_/           \
                             \
                              \_TIGGER_
                              /
       _TIGGER__             /
                \           /
                 \_TIGGER__/
_WIG__           /
      \_WIG_____/
_ZIG__/
*/

var tournament2 = []string{"LARGE RAGE", "ZEN RAGE", "RAGE"}

/* Tournament 2
        _ZEN__
              \
               \_RAGE_
_LARGE_        /
       \_RAGE_/
_RAGE__/
*/

var tournament3 = []string{"ANT BOA", "COW DUCK", "EEL FROG", "GOOSE HEN", "IGUANA JACKEL", "KITE LLAMA", "MOUSE NIT", "OCTOPUS PIG",
	"BOA COW", "FROG GOOSE", "IGUANA KITE", "MOUSE OCTOPUS", "COW GOOSE", "IGUANA OCTOPUS", "GOOSE OCTOPUS", "OCTOPUS"}

/* Tournament 3
_ANT_____
         \_BOA_____
_BOA_____/         \
                    \_COW_____
_COW_____           /         \
         \_COW_____/           \
_DUCK____/                      \
                                 \_GOOSE___
_EEL_____                        /         \
         \_FROG____             /           \
_FROG____/         \           /             \
                    \_GOOSE___/               \
_GOOSE___           /                          \
         \_GOOSE___/                            \
_HEN_____/                                       \
                                                  \_OCTOPUS_
_IGUANA__                                         /
         \_IGUANA__                              /
_JACKEL__/         \                            /
                    \_IGUANA__                 /
_KITE____           /         \               /
         \_KITE____/           \             /
_LLAMA___/                      \           /
                                 \_OCTOPUS_/
_MOUSE___                        /
         \_MOUSE___             /
_NIT_____/         \           /
                    \_OCTOPUS_/
_OCTOPUS_           /
         \_OCTOPUS_/
_PIG_____/
*/

var tournament4 = []string{"IGUANA JACKEL", "KITE LLAMA", "MOUSE NIT", "OCTOPUS PIG", "IGUANA KITE", "MOUSE OCTOPUS", "IGUANA OCTOPUS", "GOOSE OCTOPUS", "OCTOPUS"}

/* Tournament 4
                                  _GOOSE___
                                           \
                                            \
                                             \
                                              \
                                               \
                                                \
                                                 \
                                            	  \_OCTOPUS_
_IGUANA__                                         /
         \_IGUANA__                              /
_JACKEL__/         \                            /
                    \_IGUANA__                 /
_KITE____           /         \               /
         \_KITE____/           \             /
_LLAMA___/                      \           /
                                 \_OCTOPUS_/
_MOUSE___                        /
         \_MOUSE___             /
_NIT_____/         \           /
                    \_OCTOPUS_/
_OCTOPUS_           /
         \_OCTOPUS_/
_PIG_____/
*/

var tournament5 = []string{"IGUANA JACKEL", "KITE LLAMA", "IGUANA KITE", "IGUANA OCTOPUS", "GOOSE OCTOPUS", "OCTOPUS"}

/* Tournament 5
                                  _GOOSE___
                                           \
                                            \
                                             \
                                              \
                                               \
                                                \
                                                 \
                                                  \_OCTOPUS_
_IGUANA__                                         /
         \_IGUANA__                              /
_JACKEL__/         \                            /
                    \_IGUANA__                 /
_KITE____           /         \               /
         \_KITE____/           \             /
_LLAMA___/                      \           /
                                 \_OCTOPUS_/
                                 /
                                /
                               /
                     _OCTOPUS_/
*/

var tournament6 = []string{"KITE LLAMA", "IGUANA KITE", "IGUANA OCTOPUS", "GOOSE OCTOPUS", "OCTOPUS"}

/* Tournament 6
                               _GOOSE___
                                        \
                                         \
                                          \
                                           \
                                            \
                                             \
                                              \
                                               \_OCTOPUS_
                                               /
        _IGUANA_                              /
                \                            /
                 \_IGUANA__                 /
_KITE__          /         \               /
       \_KITE___/           \             /
_LLAMA_/                     \           /
                              \_OCTOPUS_/
                              /
                             /
                            /
                  _OCTOPUS_/
*/

// this java code http://www.icpc-midcentral.us/archives/2003/tourn/tourn.java fails for this tournament 6 (see below)
// the reason is that the number of level can't really be solved with Log2 (or Log/Log(2)) because too many parties
// can be missing.
// the java code outputs:
/*
                     _GOOSE___
                              \
                               \
                                \
                                 \_OCTOPUS_
  _IGUANA                        /
         \_IGUANA__             /
_LLAMAE__/         \           /
                    \_OCTOPUS_/
                    /
          _OCTOPUS_/
*/
// failed

var tournament7 = []string{"COW DUCK", "GNAT HEN", "BAT COW", "FROG GNAT", "IGUANA KITE", "MOUSE ORC", "COW GNAT", "IGUANA ORC", "GNAT ORC", "ORC"}

/*
       _BAT____
               \
                \_COW____
_COW__          /        \
      \_COW____/          \
_DUCK_/                    \
                            \_GNAT_
                            /      \
       _FROG___            /        \
               \          /          \
                \_GNAT___/            \
_GNAT_          /                      \
      \_GNAT___/                        \
_HEN__/                                  \
                                          \_ORC_
                                          /
       _IGUANA_                          /
               \                        /
                \_IGUANA_              /
                /        \            /
       _KITE___/          \          /
                           \        /
                            \_ORC__/
                            /
       _MOUSE__            /
               \          /
                \_ORC____/
                /
       _ORC____/
*/

var tournament8 = []string{"GNAT HEN", "BAT COW", "FROG GNAT", "IGUANA KITE", "MOUSE ORC", "COW GNAT", "IGUANA ORC", "GNAT ORC", "ORC"}

/*
       _BAT____
               \
                \_COW____
                /        \
       _COW____/          \
                           \
                            \_GNAT_
                            /      \
       _FROG___            /        \
               \          /          \
                \_GNAT___/            \
_GNAT_          /                      \
      \_GNAT___/                        \
_HEN__/                                  \
                                          \_ORC_
                                          /
       _IGUANA_                          /
               \                        /
                \_IGUANA_              /
                /        \            /
       _KITE___/          \          /
                           \        /
                            \_ORC__/
                            /
       _MOUSE__            /
               \          /
                \_ORC____/
                /
       _ORC____/
*/

var tournament9 = []string{"DIG BIG", "WIG ZIG", "FIG BIG", "TIG WIG", "FIG TIG", "TIG"}

/*
      _FIG_
           \
            \_FIG_
_DIG_       /     \
     \_BIG_/       \
_BIG_/              \
                     \_TIG_
                     /
      _TIG_         /
           \       /
            \_TIG_/
_WIG_       /
     \_WIG_/
_ZIG_/
*/

var tournament10 = []string{"TIG DIG", "WIG ZIG", "FIG BIG", "TIG WIG", "FIG TIG", "TIG"}

/*
      _FIG_
           \
            \_FIG_
            /     \
      _BIG_/       \
                    \
                     \_TIG_
_TIG_                /
     \_TIG_         /
_DIG_/     \       /
            \_TIG_/
_WIG_       /
     \_WIG_/
_ZIG_/
*/

var tournament11 = []string{"HHHHHHH BBB", "BBB"}

/*
_HHHHHHH_
         \_BBB_
_BBB_____/
*/

var tournament12 = []string{"HHHHHHH BBB", "HHHHHHH"}

/*
_HHHHHHH_
         \_HHHHHHH_
_BBB_____/
*/

var tournament13 = []string{"HAG HAM", "HAG HAS", "HAS"}

/*
_HAG_
     \_HAG_
_HAM_/     \
            \_HAS_
            /
      _HAS_/
*/

var tournaments = [][]string{tournament1, tournament2, tournament3, tournament4, tournament5, tournament6,
	tournament7, tournament8, tournament9, tournament10, tournament11, tournament12, tournament13}

// mutates the original slice
func reverse(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// math.Max is costly
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Rule: the name of a player must be unique for a given turn
func teamsForTurn(selection []string) (teams []string, maxNameSize int) {
	player := map[string]int{}

	for _, team := range selection {

		t := strings.Split(team, " ")
		var tempSize int

		for i, p := range t {
			var nameSize int
			if _, ok := player[p]; ok {
				// if player exists: complete turn, return
				return
			}
			player[p] = 1

			// determine maxNameSize
			// switch guard: don't take length of t[0] in account if t[1] is not valid
			switch {
			// winner case (only one player), don't wait for t[1]
			case len(t) == 1:
				nameSize = len(p)
			// memoize length of t[0] waiting for validation of t[1]
			case i == 0:
				tempSize = len(p)
			// t[1] valid, take max length of t[0] and t[1]
			default:
				nameSize = max(tempSize, len(p))
			}
			maxNameSize = max(maxNameSize, nameSize)
		}
		teams = append(teams, team)
	}

	return
}

func teamsByTurn(teams []string) (teamsByTurn [][]string, maxNameSizeByTurn []int) {
	// we want to treat the teams backwards (tree)
	reverse(teams)

	for i := 0; i < len(teams); {
		// returns teams with unique name for players = a turn
		turnTeams, turnMaxNameSize := teamsForTurn(teams[i : i+len(teams[i:])])

		// we want the teams backwards for a turn (faster retrieval with findPlayerInTurn)
		reverse(turnTeams)

		teamsByTurn = append(teamsByTurn, turnTeams)
		maxNameSizeByTurn = append(maxNameSizeByTurn, turnMaxNameSize)

		// next turn
		i += len(turnTeams)
	}

	// reverse maxNameSizeByTurn
	for i, j := 0, len(maxNameSizeByTurn)-1; i < j; i, j = i+1, j-1 {
		maxNameSizeByTurn[i], maxNameSizeByTurn[j] = maxNameSizeByTurn[j], maxNameSizeByTurn[i]
	}

	return
}

// which column to write for each turn
func columns(sizeByTurn []int) []int {
	cols := []int{0}
	for i := 1; i < len(sizeByTurn); i++ {
		cols = append(cols, cols[i-1]+sizeByTurn[i-1]+2+1<<uint64(i-1))
	}
	return cols
}

// Tournament embeds containers
type Tournament struct {
	turns     [][]string // turns (beginning with the winner)
	nameSizes []int      // max name size by turn
	columns   []int      // column where to write a player by turn
	grid      [][]string // grid to print
}

// NewTournament returns a tournament
func NewTournament(teams []string) *Tournament {
	turns, nameSizes := teamsByTurn(teams)
	columns := columns(nameSizes)
	rows := 1<<uint(len(turns)-1)*2 - 1
	cols := columns[len(turns)-1] + nameSizes[len(nameSizes)-1] + 2
	grid := make([][]string, rows)
	for i := range grid {
		grid[i] = make([]string, cols)
		for j := range grid[i] {
			grid[i][j] = " "
		}
	}

	return &Tournament{
		turns:     turns,
		nameSizes: nameSizes,
		columns:   columns,
		grid:      grid,
	}
}

// Print outputs the graph from grid (after having browsed the tree)
func (t Tournament) Print() {
	if len(t.turns) > 0 {
		// before printing, go to browse the tree and fill the grid
		initialRow := 1<<uint(len(t.turns)-1) - 1
		t.browseTree(t.turns[0][0], 0, initialRow)
		// we are coming back: print the winner in grid
		t.printPlayer(t.turns[0][0], "", len(t.turns)-1, initialRow)
		// render the grid
		for i := range t.grid {
			p := strings.TrimRight(strings.Join(t.grid[i], ""), " ")
			// get rid of all blank lines
			if p != "" {
				fmt.Println(p)
			}
		}
	}
}

// print the player with diagonal in grid
func (t Tournament) printPlayer(player, diagonal string, index, row int) {
	// print player
	p := strings.Split(fmt.Sprintf("_%s%s", player, strings.Repeat("_", t.nameSizes[index]+1-len(player))), "")
	for i, c := range p {
		t.grid[row][t.columns[index]+i] = c
	}

	// print diagonal
	for i := 0; i < 1<<uint(index); i++ {
		switch diagonal {
		case "\\":
			t.grid[row+1+i][t.columns[index]+len(p)+i] = diagonal
		case "/":
			t.grid[row-i][t.columns[index]+len(p)+i] = diagonal
		}
	}
}

// search for a player in a turn
func (t Tournament) findPlayerInTurn(iTurn int, player string) (team []string) {
	if iTurn > len(t.turns)-1 {
		return
	}
	for _, pair := range t.turns[iTurn] {
		p := strings.Split(pair, " ")
		if len(p) == 2 && (player == p[0] || player == p[1]) {
			return p
		}
	}
	return
}

// dfs: sort of...
func (t Tournament) browseTree(player string, iTurn, row int) {
	iTurn++

	team := t.findPlayerInTurn(iTurn, player)
	if team == nil {
		return
	}

	up, down := team[0], team[1]

	// what is index ?
	// index is the reverse iTurn.
	// eg: working backwards, winner is at indice 0 but at the end on the original array.
	// if we have one array and the same array reverse, index is the indice from the end
	// of the original array (cf example in 'offset' comment)
	index := len(t.turns) - 1 - iTurn

	// what is offset ?
	// offset is 2^index (value to add or subtract to previous row)
	// rule is: row =  previous row -/+ offset
	// eg:
	// t.turns:         [ [ "RAGE" ] [  "ZEN    RAGE" ] [  "HAM   ZEN",  "LARGE  RAGE" ] ]
	// iTurn:                 0              1                        2
	// index:                 2              1                        0
	// initial row            2^(len(t.turns)-1) = 4
	// -/+ offset:                      -2^1    +2^1      -2^0   +2^0     -2^0   +2^0
	// row                    4         4-2=2   4+2=6     2-1=1  2+1=3    6-1=5  6+1=7
	// t.turns:         [ [ "RAGE" ] [  "ZEN    RAGE" ] [  "HAM   ZEN",  "LARGE  RAGE" ] ]
	// eg: graph to print:
	// 1 _HAM___
	// 2        \_ZEN__
	// 3 _ZEN___/      \
	// 4                \_RAGE_
	// 5 _LARGE_        /
	// 6        \_RAGE_/
	// 7 _RAGE__/
	// so, we know on which row to write a player (columns are already in the struct)
	offset := int(1 << uint(index))

	t.browseTree(up, iTurn, row-offset)
	t.printPlayer(up, "\\", index, row-offset)

	t.browseTree(down, iTurn, row+offset)
	t.printPlayer(down, "/", index, row+offset)
}

func main() {
	for i, tournament := range tournaments {
		fmt.Printf("Tournament %d\n", i+1)
		NewTournament(tournament).Print()
	}
}
