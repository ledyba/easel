#include <stdio.h>
#include <math.h>

static const int lobe = 11;
const float PI = 3.14159265358979323846264338327950288;
#define TYPE float
#define SIN sinf

double L(TYPE x) {
	if (fabsf((float)x) <= 0.00000001) {
		return 1;
	}
	TYPE px = PI * x;
	double r = (double)(lobe) / ((double)(px) * (double)(px));
	r *= SIN(px);
	r *= SIN(px / (TYPE)(lobe));
	return r;
}

int main() {
  double sum = 0;
  double deltaX = 0.12;
  double deltaY = 0.12;
  for(int dx = -(lobe-1); dx <= lobe; dx++) {
		for(int dy = -(lobe-1); dy <= lobe; dy++) {
			double const w = (double)(L(dx + deltaX)) * (double)(L(dy + deltaY));
      sum += w;
      printf("% 4.3f ", w);
		}
    printf("\n");
	}
  printf("Total: %f\n", sum);
  return 0;
}
