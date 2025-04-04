# The Ramanujan Machine

**NOTE: This is NOT being actively developed and not ready for use.**

This is an implementation of the [Ramanujan Machine](http://www.ramanujanmachine.com/) ideas as described in the paper:

> [The Ramanujan Machine: Automatically Generated Conjectures on Fundamental Constants](https://arxiv.org/pdf/1907.00205.pdf)

The Ramanujan Machine is a brute force, algorithmic approach to finding conjectures
for [fundamental constants](https://www.wikiwand.com/en/List_of_mathematical_constants) 
of mathematics. The original Python implementation has already discovered [dozens 
of new conjectures](http://www.ramanujanmachine.com/results/).

The goal of this project is to extend the ideas of the original project by adding additional algorithms and constants but primarily to create a highly scalable system that can harness the vast computing horsepower of the cloud. I am using Go because of its ability to maximize hardware utilization but still retain an easy to read, elegant implementation.

This program works by calculating a value for the left and right hand sides of an equation. Each side uses some algorithm to calculate a value. Classically, the Ramanujan Machine uses a continued fraction on the right hand side and a rational function on the left but you aren't limited to those. For example I am also using a nested radical as an option.

The arguments for each algorithm, are generated using some kind of sequence generator, pre-calculated values or even well known constants. For example, a simple sequence is an array of whole integers such as `[0,1,2,3,4,5...]`. The values from each sequence are used as coefficients for polynomials or numerator / denoninator pairs for instance.

Lots and lots of sequences are generated, algorithms are solved using all combinations of the sequence values and the result is hashed and stored.

After all that, the hashes are compared for the left-hand and right-hand sides for equality. If the hashes are equal, then the algorithms + arguments for each side are equal, at least to 64-bits of precision. If the results are non-trivial (i.e. 2 / 1 = 2 * 1) then we might have found a novel conjecture.  If the results are already known, we flag it in the database so it isn't output again. We can then recalculate the equation using higher levels of precision.  

To generate additional results, post processing functions can be applied to the original results such as taking the square root of the result, cubed, inverse (1/result), exponential values (result**n), squared inverse (1/result^2) etc.

### Building
Simply run `make all` to build all the binaries and copy them somewhere on your path (e.g. /usr/local/bin)

### Testing
`go test ./...`

### Running

### Attribution
Social image "Raspberry Pi" by craigmdennis is marked under CC0 1.0. To view the terms, visit https://creativecommons.org/licenses/cc0/1.0/
