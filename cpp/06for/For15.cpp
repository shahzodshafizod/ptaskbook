#include <iostream>
using namespace std;

int main()
{
	//Task("For15");
	double A;
	cin >> A;
	int N;
	cin >> N;
	double degreeA = 1;
	for (int i = 1; i <= N; i++)
		degreeA *= A;
	cout << degreeA;
	
	return 0;
}
