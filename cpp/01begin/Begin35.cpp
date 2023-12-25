#include <iostream>
using namespace std;

int main()
{
	//Task("Begin35");
	double V, U, T1, T2;
	cin >> V >> U >> T1 >> T2;
	double SRiver = (V - U) * T2;
	double SLake = V * T1;
	cout << SRiver + SLake;
	
	return 0;
}
