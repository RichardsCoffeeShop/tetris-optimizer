# tetris-optimizer 🗂️

This is a simple script that reads a file with a list of tetris pieces and tries to find the best way to place them in a grid.

The algorithm is based on a recursive backtracking approach, which tries all possible combinations of pieces in smaller grids and then expands the grid if tried all possible combinations.

The examples:

```console
$ go run main.go examples/easy.txt

🔍 Started solving...
⌛ Time took to solve:  102ms
📝 Result:

ABBBB
ACCC.
A..C.
ADD..
DD...
```

## How to run

To run the script, use `go run main.go <input_file>`

It has automated tests and can be run with `go test` command.

###### [⚠️] Veryhard example could take a while to solve since the algorithm is required to try all possible combinations.
