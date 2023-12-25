#include <iostream>
using namespace std;

int main()
{
	//Task("String3");
	char C;
	cin >> C;
	int code = int(C);
	cout << char(code-1) << char(code+1);
	
	return 0;
}
