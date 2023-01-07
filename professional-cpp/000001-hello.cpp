import <iostream>;

using namespace std;

int main(int argc, char const *argv[]) {
  cout << "Hello, World!" << endl;
  return 0;
}

/*
eecd1975f4f3:~# g++ -std=c++20 -fmodules-ts -Wall -o hello.exe  hello.cpp
In module imported at hello.cpp:1:1:
/usr/include/c++/12.2.1/iostream: error: failed to read compiled module: No such
file or directory /usr/include/c++/12.2.1/iostream: note: compiled module file
is 'gcm.cache/./usr/include/c++/12.2.1/iostream.gcm'
/usr/include/c++/12.2.1/iostream: note: imports must be built before being
imported /usr/include/c++/12.2.1/iostream: fatal error: returning to the gate
for a mechanical issue compilation terminated.
*/
// g++ -std=c++20 -fmodules-ts -xc++-system-header iostream
// g++ -std=c++20 -fmodules-ts -Wall -o hello.exe  hello.cpp
// ./hello.exe