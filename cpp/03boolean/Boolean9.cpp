#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean9");
	int A, B;
	cin >> A >> B;
	bool uneven = (A % 2 != 0) || (B % 2 != 0);
	cout << uneven;
	
	return 0;
}
