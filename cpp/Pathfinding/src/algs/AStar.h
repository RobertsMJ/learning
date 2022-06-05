#ifndef __A_STAR_H
#define __A_STAR_H
#include <unordered_map>

#include "../utils/PriorityQueue.h"

template <typename Location, typename cost_t>
inline cost_t heuristic(const Location& a, const Location& b) {
  return std::abs(a.x - b.x) + std::abs(a.y - b.y);
}

template <typename Graph>
void a_star_search(const Graph& graph, const typename Graph::Location& start,
                   const typename Graph::Location& goal,
                   std::unordered_map<typename Graph::Location,
                                      typename Graph::Location>& came_from,
                   std::unordered_map<typename Graph::Location,
                                      typename Graph::Cost>& cost_so_far) {
  typedef typename Graph::Location Location;
  typedef typename Graph::Cost Cost;

  PriorityQueue<Location, Cost> frontier;

  frontier.put(start, 0);

  came_from[start] = start;
  cost_so_far[start] = 0;

  while (!frontier.empty()) {
    auto current = frontier.get();

    if (current == goal) break;

    for (Location next : graph.neighbors(current)) {
      Cost new_cost = cost_so_far[current] + graph.cost(current, next);

      if (cost_so_far.find(next) == cost_so_far.end() ||
          new_cost < cost_so_far[next]) {
        cost_so_far[next] = new_cost;
        Cost priority = new_cost + heuristic<Location, Cost>(next, goal);
        frontier.put(next, priority);
        came_from[next] = current;
      }
    }
  }
}

#endif