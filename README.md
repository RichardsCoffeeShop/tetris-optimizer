# tetris-optimizer 🗂️

This is a simple script that reads a file with a list of tetris pieces and tries to find the best way to place them in a grid.

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

#### ⚠️: Hard exam could take a while to solve since the algorithm is required to try all possible combinations.
