#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean6");
	int A, B, C;
	cin >> A >> B >> C;
	bool progress = (A < B) && (B < C);
	cout << progress;
	
	return 0;
}
