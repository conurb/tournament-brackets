const tournament1 = ["BIG DIGS", "WIG ZIG", "BIG FIGURES", "TIGGER WIG", "FIGURES TIGGER", "TIGGER"]
const tournament2 = ["LARGE RAGE", "ZEN RAGE", "RAGE"]
const tournament3 = ["ANT BOA", "COW DUCK", "EEL FROG", "GOOSE HEN", "IGUANA JACKEL", "KITE LLAMA", "MOUSE NIT", "OCTOPUS PIG", "BOA COW", "FROG GOOSE", "IGUANA KITE", "MOUSE OCTOPUS", "COW GOOSE", "IGUANA OCTOPUS", "GOOSE OCTOPUS", "OCTOPUS"]
const tournament4 = ["IGUANA JACKEL", "KITE LLAMA", "MOUSE NIT", "OCTOPUS PIG", "IGUANA KITE", "MOUSE OCTOPUS", "IGUANA OCTOPUS", "GOOSE OCTOPUS", "OCTOPUS"]
const tournament5 = ["IGUANA JACKEL", "KITE LLAMA", "IGUANA KITE", "IGUANA OCTOPUS", "GOOSE OCTOPUS", "OCTOPUS"]
const tournament6 = ["KITE LLAMA", "IGUANA KITE", "IGUANA OCTOPUS", "GOOSE OCTOPUS", "OCTOPUS"]
const tournament7 = ["COW DUCK", "GNAT HEN", "BAT COW", "FROG GNAT", "IGUANA KITE", "MOUSE ORC", "COW GNAT", "IGUANA ORC", "GNAT ORC", "ORC"]

const tournaments = [tournament1, tournament2, tournament3, tournament4, tournament5, tournament6, tournament7]

function newTournament(tournament) {

    function max(obj) {
        let max = 0
        for (let o in obj) {
            if (obj[o] > max) max = obj[o]
        }
        return max
    }

    function makeTurns(ar) {
        let turns = [], nameSizeByTurn = [], names = {}, j = 0
        ar.forEach((team, i) => {
            const [p0, p1] = team.split(' ')
            if (names[p0] || names[p1]) {
                turns.push(ar.slice(j, i).reverse())
                nameSizeByTurn.push(max(names))
                names = {}; j = i
            }
            names[p0] = p0.length
            if (p1) names[p1] = p1.length
        })
        turns.push(ar.slice(j).reverse())
        nameSizeByTurn.push(max(names))
        return { turns, nameSizeByTurn: nameSizeByTurn.reverse() }
    }

    function colByTurn(ar) {
        const cols = [0]
        for (let i = 1; i < ar.length; i++) {
            cols.push(cols[i - 1] + ar[i - 1] + 2 + (1 << (i - 1)))
        }
        return cols
    }

    function printPlayerInGrid(player, diagonal, index, row) {
        const p = `_${player}${'_'.repeat(nameSizeByTurn[index] - player.length + 1)}`.split('')
        p.forEach((c, i) => {
            grid[row][cols[index] + i] = c
        })
        for (let i = 0; i < (1 << index); i++) {
            switch (diagonal) {
                case '\\': grid[row + 1 + i][cols[index] + p.length + i] = diagonal; break;
                case '/': grid[row - i][cols[index] + p.length + i] = diagonal
            }
        }
    }

    function browseTree(player, iTurn, row) {
        iTurn++
        if (iTurn > turns.length - 1) return

        const team = turns[iTurn].find(team => {
            const [p0, p1] = team.split(' ')
            return p0 === player || p1 === player
        })

        if (!team) return

        const [up, down] = team.split(' ')
        const index = turns.length - 1 - iTurn
        const offset = 1 << index

        browseTree(up, iTurn, row - offset)
        printPlayerInGrid(up, '\\', index, row - offset)

        browseTree(down, iTurn, row + offset)
        printPlayerInGrid(down, '/', index, row + offset)
    }

    const { turns, nameSizeByTurn } = makeTurns(tournament.reverse())
    const cols = colByTurn(nameSizeByTurn)
    const grid = new Array((1 << turns.length - 1) * 2 - 1)
    const rowSize = cols[cols.length - 1] + nameSizeByTurn[nameSizeByTurn.length - 1] + 2
    for (let i = 0; i < grid.length; i++) {
        grid[i] = new Array(rowSize).fill(' ')
    }

    return {
        print() {
            const row = (1 << turns.length - 1) - 1
            browseTree(turns[0][0], 0, row)
            printPlayerInGrid(turns[0][0], '', turns.length - 1, row)
            grid.forEach(row => {
                const line = row.join('').replace(/\s+$/, '')
                if (line != '') console.log(line)
            })
        }
    }
}

tournaments.forEach((tournament, i) => {
    console.log(`Tournament ${i + 1}`)
    newTournament(tournament).print()
})