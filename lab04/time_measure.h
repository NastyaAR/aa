#ifndef __TIME_MEASURE_H__
#define __TIME_MEASURE_H__

#include <chrono>
#include <vector>

#define N_REPS 1000
#define LOGICAL_CORS 12

typedef struct time_info
{
    std::vector<int> sizes;
    std::vector<double> seconds;
} time_info_t;

std::vector<time_info_t> measure_time();

#endif