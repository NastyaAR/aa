#include "time_measure.h"

std::vector<time_info_t> measure_time()
{
    std::vector<time_info_t> time_info;
    auto begin = std::chrono::high_resolution_clock::now();
	auto end = std::chrono::high_resolution_clock::now();
	double elapsed_s = 0;
	double sum = 0;

    for (size_t nThreads = 2; nThreads < 4 * LOGICAL_CORS; nThreads += 2) {
        time_info_t cur_time_info;
        for (size_t i = 0; i < N_REPS; i += 50) {
            begin = std::chrono::high_resolution_clock::now();
            //какой-то код...
            end = std::chrono::high_resolution_clock::now();
            elapsed_s = (double) std::chrono::duration_cast<std::chrono::seconds>(end - begin).count();
		    sum += elapsed_s;
        }
        cur_time_info.sizes.push_back(0); //размер
        cur_time_info.seconds.push_back(sum / N_REPS);
        time_info.push_back(cur_time_info);
    }

    return time_info;
}