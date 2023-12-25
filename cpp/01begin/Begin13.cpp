#include <iostream>
using namespace std;

int main()
{
	//Task("Begin13");
	float R1, R2;
	cin >> R1 >> R2;
	const float PI = 3.14;
	float S1 = PI * R1*R1;
	float S2 = PI * R2*R2;
	cout << S1 << S2 << S1-S2;
	
	return 0;
}
