#include <iostream>
using namespace std;

int main()
{
	//Task("Integer29");
	int A, B, C;
	cin >> A >> B >> C;
	int kvads = (A / C) * (B / C);
	int freeSpace = A * B - kvads * C * C;
	cout << kvads << freeSpace;
	
	return 0;
}
