#ifndef __INP_OUT_H__
#define __INP_OUT_H__

#include <stdio.h>
#include <iostream>
#include <fstream>
#include "err.h"
#include "time_measure.h"

#define MAX_COMMAND 3

size_t inputChoice(void);
void outputMenu(void);
void formCsvs(std::vector<time_info_t> time_info);

#endif