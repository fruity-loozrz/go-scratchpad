package main

import (
	"fmt"

	ast "github.com/fruity-loozrz/go-scratchpad/internal/automationast"
	"github.com/fruity-loozrz/go-scratchpad/internal/automationparser/lexer"
	"github.com/fruity-loozrz/go-scratchpad/internal/automationparser/parser"
)

func main() {
	task := `bpm  120
5
+
+1
+1 1
+ 1
+2
+2 2
+1/2
+1/2 1/2
-
-1
-1 1
-1/2
-1/2
`

	par := parser.NewParser()
	lex := lexer.NewLexer([]byte(task))

	something, err := par.Parse(lex)
	if err != nil {
		panic(err)
	}

	prog := something.(ast.Program)
	// fmt.Printf("%#v\n", prog)

	for _, line := range prog.Lines {
		fmt.Printf("%#v\n", line)
	}
}

// 	task := `bpm  120	# set 1 beat = 60/120 s
// 5			# init head at position 5 and wait 1 beat
// +			# advance head by 1 beat and wait 1 beat
// +1			# the same as above
// +1 1		# the same as above
// + 1			# the same as above
// +2			# advance head by 2 beats and wait 1 beat
// +2 2		# advance head by 2 beats and wait 2 beats
// +1/2		# advance head by 1/2 beat and wait 1 beat
// +1/2 1/2	# advance head by 1/2 beat and wait 1/2 beat
// -			# move head back by 1 beat and wait 1 beat
// -1			# the same as above
// -1 1		# the same as above
// -1/2		# move head back by 1/2 beat and wait 1 beat
// -1/2		# move head back by 1/2 beat and stop (it never has a duration)`
