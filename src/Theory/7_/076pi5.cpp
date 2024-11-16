#include <iostream>
#include <omp.h>

using namespace std;

double calcPi(long long n) {
    double step = 1. / double(n);
    double x;
    double sum = 0.;
#pragma omp parallel for private(x) reduction(+:sum)
    for (int i = 0; i < n; ++i) {
        x = double(i) * step;
        sum += 4. / (1. + x*x);
    }

    return sum / double(n);
}

int main() {
    int n = 1e8;

    double start = omp_get_wtime();
    double pi = calcPi(n);
    double dur = omp_get_wtime() - start;

    cout << "Pi: " << pi << endl;
    cout << "Time: " << dur << endl;

    return 0;
}
