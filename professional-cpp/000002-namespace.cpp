#include <iostream>

namespace mycode
{
    using namespace std;
    void foo()
    {
        cout << "namespace: mycode, function: foo" << endl;
    }
} // namespace mycode

int main(int argc, char const *argv[])
{
    mycode::foo();
    return 0;
}
