#include <iostream>
using namespace std;

int main()
{
	//Task("While1");
	double A, B;
	cin >> A >> B;
	while (A >= B)
		A -= B;
	cout << A;
	
	return 0;
}
