#ifndef __GRID_UTILS_H
#define __GRID_UTILS_H
#include <iomanip>
#include <iostream>
#include <unordered_map>
#include <vector>

#include "GridLocation.h"
// This outputs a grid. Pass in a distances map if you want to print
// the distances, or pass in a point_to map if you want to print
// arrows that point to the parent location, or pass in a path vector
// if you want to draw the path.
template <class Graph>
void draw_grid(const Graph& graph,
               unordered_map<GridLocation, double>* distances = nullptr,
               unordered_map<GridLocation, GridLocation>* point_to = nullptr,
               vector<GridLocation>* path = nullptr,
               GridLocation* start = nullptr, GridLocation* goal = nullptr) {
  const int field_width = 3;
  cout << string(field_width * graph.width, '-') << '\n';
  for (int y = 0; y != graph.height; ++y) {
    for (int x = 0; x != graph.width; ++x) {
      GridLocation loc{x, y};
      if (graph.walls.find(loc) != graph.walls.end()) {
        cout << string(field_width, '#');
      } else if (start && loc == *start) {
        cout << " A ";
      } else if (goal && loc == *goal) {
        cout << " Z ";
      } else if (path != nullptr &&
                 find(path->begin(), path->end(), loc) != path->end()) {
        cout << " @ ";
      } else if (point_to != nullptr && point_to->count(loc)) {
        GridLocation next = (*point_to)[loc];
        if (next.x == x + 1) {
          cout << " → ";
        } else if (next.x == x - 1) {
          cout << " ← ";
        } else if (next.y == y + 1) {
          cout << " ↓ ";
        } else if (next.y == y - 1) {
          cout << " ↑ ";
        } else {
          cout << " * ";
        }
      } else if (distances != nullptr && distances->count(loc)) {
        cout << ' ' << left << setw(field_width - 1) << (*distances)[loc];
      } else {
        cout << " . ";
      }
    }
    cout << '\n';
  }
  cout << string(field_width * graph.width, '-') << '\n';
}

#endif