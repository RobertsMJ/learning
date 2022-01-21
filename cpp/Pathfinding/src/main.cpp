#include <iostream>
using namespace std;

#include "algs/BreadthFirstSearch.h"
#include "graphs/GridUtils.h"
#include "graphs/SimpleGraph.h"
#include "graphs/SquareGrid.h"

int main() {
  SquareGrid grid = make_diagram1();
  GridLocation start{7, 8}, goal{17, 2};
  auto parents = breadth_first_search(grid, start, goal);
  draw_grid(grid, nullptr, &parents, nullptr, &start, &goal);
  return 0;
}
