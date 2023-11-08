#include <iostream>
#include "inp_out.h"

#define N_REPS 1000

int main() {
	size_t choice = 10;

	outputMenu();

	while (choice) {
		try {
			choice = inputChoice();
		}
		catch (const input_error& err)
    	{
        	std::cout << err.what() << std::endl;
			continue;
    	}

		switch (choice)
		{
		case 0:

			break;
		case 1:
			std::cout << "1\n" << std::endl;
			break;
		case 2:
			std::cout << "2\n" << std::endl;
			break;
		case 3:
			outputMenu();
			break;
		}
	}

	return 0;
} 
