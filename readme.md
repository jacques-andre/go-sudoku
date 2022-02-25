![Go build](https://github.com/jacques-andre/go-sudoku/actions/workflows/go.yml/badge.svg)


## Solving Approaches:

### Brute force:

Generate **all** possible solutions, then validate **all** possible solutions.
- Slow

### Backtracking:

3 Key concepts:
- Choice
- Constraint
- Goal

**Choice:**
- Numbers 1-9

**Constraint:**
- Valid in row?
- Valid in col?
- Valid in subgrid (3x3)?

**Goal:**
- Reached Base cases:
  - Finished a row/col. > Move on


### Dancing Links


### Depth first search

....
## Resources:
- https://www.youtube.com/watch?v=JzONv5kaPJM
- https://en.wikipedia.org/wiki/NP-completeness
- https://www.youtube.com/watch?v=A80YzvNwqXA
- https://www.youtube.com/watch?v=A80YzvNwqXA
- https://en.wikipedia.org/wiki/Dancing_Links
- https://www.ocf.berkeley.edu/~jchu/publicportal/sudoku/0011047.pdf
- https://www.youtube.com/watch?v=A80YzvNwqXA 
