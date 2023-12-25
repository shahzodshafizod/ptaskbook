#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean8");
	int A, B;
	cin >> A >> B;
	bool unevens = (A % 2 != 0) && (B % 2 != 0);
	cout << unevens;
	
	return 0;
}
