// Given arrays of weights, values, and capacity

// recursively search for the maximum value attainable if item n is included
// store the result of including item m with remaining capacity k
#include <vector>
using std::vector;
#include <map>
using std::map;
#include <string>
using std::string;
using std::to_string;
#include <limits>
using std::numeric_limits;
#include <iostream>
using std::cout;
using std::endl;
#include <algorithm>
using std::max;

int knapsack(const vector<int>& values, const vector<int>& weights,
             int capacity, int item, map<string, int>& subproblems);
int main() {
  vector<int> v{10, 20, 30, 40};
  vector<int> w{30, 10, 40, 20};
  int capacity = 40;
  map<string, int> subproblems;

  cout << knapsack(v, w, capacity, v.size() - 1, subproblems) << endl;
}

int knapsack(const vector<int>& values, const vector<int>& weights,
             int capacity, int item, map<string, int>& subproblems) {
  // Full
  if (capacity < 0) return numeric_limits<int>::min();
  // full or out of items
  if (capacity == 0 || item < 0) return 0;

  string key = to_string(item) + "|" + to_string(capacity);

  // If we haven't solved this subproblem before
  if (!subproblems.contains(key)) {
    // See if it's better to include or exclude it
    // include
    auto include =
        values[item] + knapsack(values, weights, capacity - weights[item],
                                item - 1, subproblems);

    // exclude
    auto exclude = knapsack(values, weights, capacity, item - 1, subproblems);

    subproblems[key] = max(include, exclude);
  }
  return subproblems[key];
}