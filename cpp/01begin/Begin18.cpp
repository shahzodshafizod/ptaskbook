#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("Begin18");
	float A, B, C;
	cin >> A >> B >> C;
	float AC = abs(C - A);
	float BC = abs(C - B);
	cout << AC * BC;
	
	return 0;
}
