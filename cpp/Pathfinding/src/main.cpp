#include <iostream>
using namespace std;

#include "algs/AStar.h"
#include "algs/BreadthFirstSearch.h"
#include "algs/Dijkstra.h"
#include "algs/heuristics/Manhattan.h"
#include "graphs/GridLocation.h"
#include "graphs/GridUtils.h"
#include "graphs/GridWithWeights.h"
#include "graphs/SimpleGraph.h"
#include "graphs/SquareGrid.h"

// // DFS
// int main() {
//   SquareGrid grid = make_diagram1();
//   GridLocation start{7, 8}, goal{17, 2};
//   auto parents = breadth_first_search(grid, start, goal);
//   draw_grid(grid, nullptr, &parents, nullptr, &start, &goal);
//   return 0;
// }

// // Dijkstra
// int main() {
//   GridWithWeights grid = make_diagram4();
//   GridLocation start{1, 4}, goal{8, 3};
//   unordered_map<GridLocation, GridLocation> came_from;
//   unordered_map<GridLocation, double> cost_so_far;
//   dijkstra_search(grid, start, goal, came_from, cost_so_far);
//   draw_grid(grid, nullptr, &came_from, nullptr, &start, &goal);
//   cout << endl;
//   vector<GridLocation> path = reconstruct_path(start, goal, came_from);
//   draw_grid(grid, nullptr, nullptr, &path, &start, &goal);
//   cout << endl;
//   draw_grid(grid, &cost_so_far, nullptr, nullptr, &start, &goal);
// }

// A*
int main() {
  using LocationT = GridLocation;
  using CostT = double;

  auto grid = make_diagram4<LocationT, CostT>();
  auto heuristic = manhattan<LocationT, CostT>;

  LocationT start{1, 9}, goal{8, 3};
  unordered_map<LocationT, LocationT> came_from;
  unordered_map<LocationT, CostT> cost_so_far;

  a_star_search(grid, heuristic, start, goal, came_from, cost_so_far);

  draw_grid(grid, nullptr, &came_from, nullptr, &start, &goal);
  cout << endl;
  vector<LocationT> path = reconstruct_path(start, goal, came_from);
  draw_grid(grid, nullptr, nullptr, &path, &start, &goal);
  cout << endl;
  draw_grid(grid, &cost_so_far, nullptr, nullptr, &start, &goal);
}