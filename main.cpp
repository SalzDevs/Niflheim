#include <iostream>
#include <cstddef>

enum Direction{
  RIGHT,
  LEFT,
  UP,
  DOWN,
};

struct coordinates{
  size_t x;
  size_t y;
};

struct gradient{
  double x;
  double y;
};

struct GridPoint{
  coordinates cords;
  gradient grad;
};

struct  Grid{
  size_t width;
  size_t height;
  GridPoint** points;
};


gradient GetGradient(Direction dir){
  switch(dir){
    case RIGHT:
      return {1.0, 0.0};
    case LEFT:
      return {-1.0, 0.0};
    case UP:
      return {0.0, 1.0};
    case DOWN:
      return {0.0, -1.0};
    default:
      return {0.0, 0.0};
  }
}


gradient GetRandomGradient(size_t x, size_t y){
  // Simple deterministic pseudo-random generator based on coordinates
  size_t seed = x * 31 + y * 17; // Combine x and y to create a seed
  Direction dir = static_cast<Direction>(seed % 4); // Get a direction based on the seed
  return GetGradient(dir);
}

void InitGrid(Grid& grid, size_t width, size_t height){
  grid.width = width;
  grid.height = height;
  grid.points = new GridPoint* [height];

  for(size_t i = 0; i < height; ++i){
    grid.points[i] = new GridPoint[width];
    for(size_t j = 0; j < width; ++j){
      grid.points[i][j].cords.x = j;
      grid.points[i][j].cords.y = i;
      grid.points[i][j].grad.x = 0.0;
      grid.points[i][j].grad.y = 0.0;
    }
  }

}



int main() {
    Grid grid;
    InitGrid(grid, 5, 5);

    for(size_t i = 0; i < grid.height; ++i){
        for(size_t j = 0; j < grid.width; ++j){
            grid.points[i][j].grad = GetRandomGradient(j, i);
        }
    }

    std::cout << "\nAfter assigning random gradients:\n";
    for(size_t i = 0; i < grid.height; ++i){
        for(size_t j = 0; j < grid.width; ++j){
            std::cout << "Grid Point (" << grid.points[i][j].cords.x << ", " << grid.points[i][j].cords.y << ") - Gradient: (" 
                      << grid.points[i][j].grad.x << ", " << grid.points[i][j].grad.y << ")\n";
        }
    }

    return 0;
}
