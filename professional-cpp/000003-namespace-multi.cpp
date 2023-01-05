#include <iostream>

namespace Libs::Network::FTP
{
    void foo()
    {
        std::cout << "Libs::Network::FTP::foo()" << std::endl;
    }
} // namespace Libs::Network::FTP

namespace MyFTP = Libs::Network::FTP;

int main(int argc, char const *argv[])
{
    MyFTP::foo();
    return 0;
}
