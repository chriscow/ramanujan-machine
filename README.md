# The Ramanujan Machine
This is an implementation of the Ramanujan Machine ideas as described in the paper:

[The Ramanujan Machine: Automatically Generated Conjectures on Fundamental Constants][1]

The Ramanujan Machine is a brute force, algorithmic approach to finding conjectures
for [fundamental constants](https://www.wikiwand.com/en/List_of_mathematical_constants) 
of mathematics. The original Python implementation has already discovered [dozens 
of new conjectures](http://www.ramanujanmachine.com/results/).

The goal of this project is to extend the ideas of the original project by adding
additional algorithms and constants but primarily to create a highly scalable system
that can harness the vast computing horsepower of the cloud.

I am using Go because of its ability to maximize hardware utilization but still 
retain an easy to read, elegant implementation.

This program works by calculating a value for the left and right hand sides of
an equation. Each side uses some algorithm to calculate a value. Classically,
the Ramanujan machine uses a continued fraction on the right hand side and a
rational function on the left.

The arguments for each algorithm, are generated using some kind of
sequence generator, pre-calculated values or even well known constants. For example, a simple sequence is an array of whole integers
such as `[0, 1, 2, 3, 4, 5 ...]`. 

The values from each sequence are used as coefficients for polynomials or numerator / 
denoninator pairs for instance.

Lots and lots of sequences are generated, algorithms are solved using all combinations
of the sequence values and the result is hashed and stored.

After all that, the hashes are compared for the left-hand and right-hand sides for
equality. If the hashes are equal, then the algorithms + arguments for each side
are obviously equal. If the results are non-trivial (i.e. 2 / 1 = 2 * 1) then
we might have found a novel conjecture.  If the results are already known, we
flag it in the database so it isn't output again. We can then recalculate the equation using higher levels of precision.  


### Building
Simply run `make all` to build all the binaries and copy them somewhere on your 
path (e.g. /usr/local/bin)

### Testing
`go test ./...`

### Running

[1]: https://arxiv.org/pdf/1907.00205.pdf [The Ramanujan Machine: Automatically Generated Conjectures on Fundamental Constants]