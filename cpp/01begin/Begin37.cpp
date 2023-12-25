#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("Begin37");
	double V1, V2, S, T;
	cin >> V1 >> V2 >> S >> T;
	double S1 = T * V1;
	double S2 = T * V2;
	cout << abs(S1 + S2 - S);
	
	return 0;
}
