package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func equalSlice2D(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i := range a {
		if !equalSlice(a[i], b[i]) {
			return false
		}
	}

	return true
}

func equalSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, s := range a {
		if s != b[i] {
			return false
		}
	}

	return true
}

// http://craigwickesser.com/2015/01/capture-stdout-in-go/
func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestTeamsByTurn(t *testing.T) {
	tt := []struct {
		tournament []string
		turns      [][]string
	}{
		{[]string{"BIG DIGS", "WIG ZIG", "BIG FIGURES", "TIGGER WIG", "FIGURES TIGGER", "TIGGER"},
			[][]string{
				[]string{"TIGGER"},
				[]string{"FIGURES TIGGER"},
				[]string{"BIG FIGURES", "TIGGER WIG"},
				[]string{"BIG DIGS", "WIG ZIG"},
			}},
		{[]string{"KITE LLAMA", "IGUANA KITE", "IGUANA OCTOPUS", "GOOSE OCTOPUS", "OCTOPUS"},
			[][]string{
				[]string{"OCTOPUS"},
				[]string{"GOOSE OCTOPUS"},
				[]string{"IGUANA OCTOPUS"},
				[]string{"IGUANA KITE"},
				[]string{"KITE LLAMA"},
			}},
		{[]string{"OCTOPUS"},
			[][]string{
				[]string{"OCTOPUS"},
			}},
	}
	for _, tc := range tt {
		turns, _ := teamsByTurn(tc.tournament)
		if !equalSlice2D(turns, tc.turns) {
			t.Errorf("\nexpected: %#v\ngot     : %#v", tc.turns, turns)
		}
	}
}

func TestOuput(t *testing.T) {
	tt := []struct {
		tournament []string
		graph      string
	}{
		{[]string{"KITE LLAMA", "IGUANA KITE", "IGUANA OCTOPUS", "GOOSE OCTOPUS", "OCTOPUS"},
			`                               _GOOSE___
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
`,
		},
		{[]string{"GNAT HEN", "BAT COW", "FROG GNAT", "IGUANA KITE", "MOUSE ORC", "COW GNAT", "IGUANA ORC", "GNAT ORC", "ORC"},
			`       _BAT____
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
`,
		},
		{[]string{"ANT BOA", "COW DUCK", "EEL FROG", "GOOSE HEN", "IGUANA JACKEL", "KITE LLAMA", "MOUSE NIT", "OCTOPUS PIG",
			"BOA COW", "FROG GOOSE", "IGUANA KITE", "MOUSE OCTOPUS", "COW GOOSE", "IGUANA OCTOPUS", "GOOSE OCTOPUS", "OCTOPUS"},
			`_ANT_____
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
`,
		},
	}
	for _, tc := range tt {
		graph := captureStdout(NewTournament(tc.tournament).Print)
		if graph != tc.graph {
			t.Errorf("expected:\n %#v\n\ngot:\n %#v\n", tc.graph, graph)
		}
	}
}
