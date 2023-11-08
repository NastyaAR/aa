#include "inp_out.h"

size_t inputChoice(void)
{
    size_t choice;
    printf("\nВведите команду: ");
    if (scanf("%zu", &choice) != 1 || choice > MAX_COMMAND)
        throw input_error("\nНекорректный ввод номера команды!");
    return choice;
}

void outputMenu(void)
{
    printf("\nМеню:");
    printf("\n1) Пункт один...;\n");
    printf("2) Пункт два...;\n");
    printf("\n3) Вывод меню;\n");
    printf("\n0) Завершение работы программы.\n");
}

static void formCsv(time_info_t time_info, char *fname, size_t nThreads)
{
    char fullName[50];
    sprintf(fullName, "./report/inc/csv/%zu/%s.csv", nThreads, fname);

    std::ofstream out;  
    out.open(fullName);
    if (out.is_open())
    {
        out << "len," << fname << std::endl;
        for (size_t i = 0; i < time_info.sizes.size(); i++) {
            out << time_info.sizes[i] << "," << time_info.seconds[i] << std::endl;
        }
    }
    out.close(); 
}

void formCsvs(std::vector<time_info_t> time_info)
{
    std::vector<char *> names = {"Без_потоков", "2_потока",
    "4_потока", "6_потоков", "8_потоков", "10_потоков"};

    for (size_t i = 0; i < names.size(); i++) {
        formCsv(time_info[i], names[i], 2 * i);
    }
}