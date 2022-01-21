#ifndef __SIMPLE_GRAPH_H
#define __SIMPLE_GRAPH_H
#include <unordered_map>
#include <vector>

struct SimpleGraph {
  unordered_map<char, vector<char>> edges;

  vector<char> neighbors(char id) { return edges[id]; }
};

#endif