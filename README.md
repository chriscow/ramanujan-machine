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

### Building
Simply run `make all` to build all the binaries and copy them somewhere on your 
path (e.g. /usr/local/bin)

### Running
[1]: https://arxiv.org/pdf/1907.00205.pdf [The Ramanujan Machine: Automatically Generated Conjectures on Fundamental Constants]