#!/bin/bash
# use chmod +x runmaster.sh in order to make it executable

g++ sequencealignment.cpp
./a.out <dataset.txt >distances_graph.txt
go run EdgeTravel.go <distances_graph.txt >results.txt
