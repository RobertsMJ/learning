#ifndef __GRID_LOCATION_H
#define __GRID_LOCATION_H

#include <cstddef>
#include <functional>
#include <tuple>

struct GridLocation {
  int x, y;
};

bool operator==(GridLocation a, GridLocation b) {
  return a.x == b.x && a.y == b.y;
}

bool operator!=(GridLocation a, GridLocation b) { return !(a == b); }

bool operator<(GridLocation a, GridLocation b) {
  return tie(a.x, a.y) < tie(b.x, b.y);
}

namespace std {
template <>
struct hash<GridLocation> {
  size_t operator()(const GridLocation& id) const noexcept {
    return hash<int>()(id.x ^ (id.y << 16));
  }
};
}  // namespace std

#endif