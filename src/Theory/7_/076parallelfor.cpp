#include <iostream>
#include <omp.h>

using namespace std;

int main() {
#pragma omp parallel for
    for (int i = 0; i < 32; ++i) {
        cout << i << endl;
    }
    return 0;
}
