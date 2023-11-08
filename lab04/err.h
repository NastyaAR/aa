#ifndef __ERR_H__
#define __ERR_H__

#include <exception>
#include <string>

class input_error: public std::exception
{
public:
    input_error(const std::string& message): message{message}
    {}
    const char* what() const noexcept override
    {
        return message.c_str();
    }
private:
    std::string message;
};

#endif