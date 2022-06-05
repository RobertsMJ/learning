#ifndef __DIJKSTRA_H
#define __DIJKSTRA_H
#include <unordered_map>
#include <vector>

#include "../utils/PriorityQueue.h"

template <typename Location, typename Graph>
void dijkstra_search(const Graph& graph, const Location& start,
                     const Location& goal,
                     std::unordered_map<Location, Location>& came_from,
                     std::unordered_map<Location, double>& cost_so_far) {
  PriorityQueue<Location, double> frontier;
  frontier.put(start, 0);

  came_from[start] = start;
  cost_so_far[start] = 0;

  while (!frontier.empty()) {
    Location current = frontier.get();

    if (current == goal) break;

    for (Location next : graph.neighbors(current)) {
      double new_cost = cost_so_far[current] + graph.cost(current, next);
      if (cost_so_far.find(next) == cost_so_far.end() ||
          new_cost < cost_so_far[next]) {
        cost_so_far[next] = new_cost;
        came_from[next] = current;
        frontier.put(next, new_cost);
      }
    }
  }
}

template <typename Location>
std::vector<Location> reconstruct_path(
    Location start, Location goal,
    std::unordered_map<Location, Location> came_from) {
  std::vector<Location> path;
  for (Location current = goal; current != start;
       current = came_from[current]) {
    path.push_back(current);
  }
  path.push_back(start);
  std::reverse(path.begin(), path.end());
  return path;
}

#endif