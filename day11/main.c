#include <stdio.h>

int power_level(int x, int y, int serial_no) {
  const int rack_id = x + 10;
  return (((rack_id * y + serial_no) * rack_id) / 100) % 10 - 5;
}

int main() {
  const int serial_no = 9110;

  int grid[300][300];

  for (int y = 0; y < 300; y++) {
    for (int x = 0; x < 300; x++) {
      grid[y][x] = power_level(x + 1, y + 1, serial_no);
    }
  }

  int mostPower = 0, mostPowerX = 0, mostPowerY = 0, mostPowerSize = 0;

  for (int size = 0; size < 300; size++) {
    for (int y = 1; y <= 300 - size; y++) {
      for (int x = 1; x <= 300 - size; x++) {
        int power = 0;

        for (int i = 0; i < size; i++) {
          for (int j = 0; j < size; j++) {
            power += grid[y + j - 1][x + i - 1];
          }
        }

        if (power >= mostPower) {
          mostPower = power;
          mostPowerX = x;
          mostPowerY = y;
          mostPowerSize = size;
        }
      }
    }

    if (size == 3) {
      printf(
          "X,Y coordinate of the top-left fuel cell of the 3x3 square with the "
          "largest total power: %d,%d\n",
          mostPowerX, mostPowerY, mostPowerSize);
    }
  }

  printf(
      "X,Y,size identifier of the square with the largest total power: "
      "%d,%d,%d\n",
      mostPowerX, mostPowerY, mostPowerSize);

  return 0;
}