#ifndef __GRID_WITH_WEIGHTS_H
#define __GRID_WITH_WEIGHTS_H

#include <unordered_set>

#include "SquareGrid.h"

template <typename location_t, typename cost_t>
struct GridWithWeights : SquareGrid<location_t> {
  using Location = location_t;
  using Cost = cost_t;

  std::unordered_set<Location> forests;
  GridWithWeights(int w, int h) : SquareGrid<Location>(w, h) {}

  cost_t cost(const Location& from_node, const Location& to_node) const {
    return forests.find(to_node) != forests.end() ? 5 : 1;
  }
};

template <typename Location, typename cost_t>
GridWithWeights<Location, cost_t> make_diagram4() {
  GridWithWeights<Location, cost_t> grid(10, 10);
  add_rect(grid, 1, 7, 4, 9);
  using L = Location;
  grid.forests = std::unordered_set<Location>{
      L{3, 4}, L{3, 5}, L{4, 1}, L{4, 2}, L{4, 3}, L{4, 4}, L{4, 5},
      L{4, 6}, L{4, 7}, L{4, 8}, L{5, 1}, L{5, 2}, L{5, 3}, L{5, 4},
      L{5, 5}, L{5, 6}, L{5, 7}, L{5, 8}, L{6, 2}, L{6, 3}, L{6, 4},
      L{6, 5}, L{6, 6}, L{6, 7}, L{7, 3}, L{7, 4}, L{7, 5}};
  return grid;
}

#endif