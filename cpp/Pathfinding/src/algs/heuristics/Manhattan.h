#ifndef __HEURISTIC_H
#define __HEURISTIC_H

template <typename location_t, typename cost_t>
inline cost_t manhattan(const location_t& a, const location_t& b) {
  return std::abs(a.x - b.x) + std::abs(a.y - b.y);
}

#endif