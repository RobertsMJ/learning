#ifndef __BREADTH_FIRST_SEARCH_H
#define __BREADTH_FIRST_SEARCH_H
#include <queue>
#include <unordered_map>

template <typename Location, typename Graph>
unordered_map<Location, Location> breadth_first_search(Graph graph,
                                                       Location start,
                                                       Location goal) {
  queue<Location> frontier;
  frontier.push(start);

  unordered_map<Location, Location> came_from;
  came_from[start] = start;

  while (!frontier.empty()) {
    auto current = frontier.front();
    frontier.pop();

    if (current == goal) break;

    for (auto next : graph.neighbors(current)) {
      if (came_from.find(next) == came_from.end()) {
        frontier.push(next);
        came_from[next] = current;
      }
    }
  }

  return came_from;
}

#endif