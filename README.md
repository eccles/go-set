# go-set

Generic set for Golang

Loosely based on the python3 sets operations using generics and iterators.

Uses a map as a synonym of a set. Members of a set must be comparable. 
Unfortunately this precludes a 'Set of sets' as maps (on which Sets are based) are not comparable.

See https://go.dev/blog/generic-interfaces for a much deeper discussion of Sets and generics.
