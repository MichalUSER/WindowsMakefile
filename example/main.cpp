#include <iostream>
#include <string>

using namespace std;

class MyClass
{
public:
    int a = 5;
    int *addr = &a;
};

int main()
{
    MyClass myObj;
    myObj.a = 6;
    cout << myObj.addr;
    cout << "\nValue of a is: " << *myObj.addr;
    MyClass *classArray[1];
    classArray[0] = &myObj;
    MyClass arrayClass = *classArray[0];
    cout << "\n";
    cout << arrayClass.a;
}