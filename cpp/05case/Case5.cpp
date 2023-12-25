#include <iostream>
using namespace std;

int main()
{
	//Task("Case5");
	int N;
	cin >> N;
	double A, B, result;
	cin >> A >> B;
	switch (N)
	{
		case 1:
			result = A + B;
			break;
		case 2:
			result = A - B;
			break;
		case 3:
			result = A * B;
			break;
		case 4:
			result = A / B;
			break;
	}
	cout << result;
	
	return 0;
}
