#include <iostream>
#include <omp.h>

#define MAXTH 8

using namespace std;

double calcPi(long long n) {
    double step = 1. / double(n);
    double sum = 0.;
    omp_set_num_threads(MAXTH);
    int numthreads;
#pragma omp parallel
{
    double x;
    double partial = 0.;
    int nth = omp_get_thread_num();
    if (nth == 0) numthreads = omp_get_num_threads();
    for (int i = nth; i < n; i += MAXTH) {
        x = double(i) * step;
        partial += 4. / (1. + x*x);
    }
#pragma omp critical
    sum += partial;
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
