#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean24");
	float A, B, C;
	cin >> A >> B >> C;
	float D = B * B - 4 * A * C;
	bool hasRoot = D >= 0;
	cout << hasRoot;
	
	return 0;
}
