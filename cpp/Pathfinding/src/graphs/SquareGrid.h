#ifndef __SQUARE_GRID_H
#define __SQUARE_GRID_H
#include <algorithm>
#include <array>
#include <unordered_set>
#include <vector>

using std::array;
using std::reverse;
using std::unordered_set;
using std::vector;

#include "./GridLocation.h"

template <typename location_t>
struct SquareGrid {
  using Location = location_t;
  static array<Location, 4> DIRS;

  int width, height;
  unordered_set<Location> walls;

  SquareGrid(int width, int height) : width(width), height(height) {}

  bool in_bounds(const Location& id) const {
    return 0 <= id.x && id.x < width && 0 <= id.y && id.y < height;
  }

  bool passable(const Location& id) const {
    return walls.find(id) == walls.end();
  }

  vector<Location> neighbors(const Location& id) const {
    vector<Location> results;

    for (auto dir : DIRS) {
      Location next{id.x + dir.x, id.y + dir.y};
      if (in_bounds(next) && passable(next)) results.push_back(next);
    }

    if ((id.x + id.y) % 2 == 0) reverse(results.begin(), results.end());

    return results;
  }
};

template <typename Location>
array<Location, 4> SquareGrid<Location>::DIRS = {
    // East, West, North, South
    GridLocation{1, 0}, GridLocation{-1, 0}, GridLocation{0, -1},
    GridLocation{0, 1}};

template <typename Location>
void add_rect(SquareGrid<Location>& grid, int x1, int y1, int x2, int y2) {
  for (int x = x1; x < x2; ++x) {
    for (int y = y1; y < y2; ++y) {
      grid.walls.insert(Location{x, y});
    }
  }
}

template <typename Location>
SquareGrid<Location> make_diagram1() {
  SquareGrid<Location> grid(30, 15);
  add_rect<Location>(grid, 3, 3, 5, 12);
  add_rect<Location>(grid, 13, 4, 15, 15);
  add_rect<Location>(grid, 21, 0, 23, 7);
  add_rect<Location>(grid, 23, 5, 26, 7);
  return grid;
}

#endif