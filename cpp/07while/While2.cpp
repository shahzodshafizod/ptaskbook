#include <iostream>
using namespace std;

int main()
{
	//Task("While2");
	double A, B;
	cin >> A >> B;
	int porchaho = 0;
	while (A >= B)
	{
		A -= B;
		porchaho++;
	}
	cout << porchaho;
	
	return 0;
}
