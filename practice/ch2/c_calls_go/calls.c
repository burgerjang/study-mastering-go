#include <stdio.h>

extern int sum(int a, int b);

int main(int argc, char**argv) {
	int r = sum(1, 2);
	printf("%d\n", r);
}
