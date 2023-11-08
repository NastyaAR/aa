#include "algs.h"

std::mutex mutex;

void getMaxInArray(int *arr, size_t len, int *globalMax)
{
    for (size_t i = 0; i < len; i++) {
        mutex.lock();
        if (*globalMax < arr[i]) {
            *globalMax = arr[i];
        }
        mutex.unlock();
    }
}

