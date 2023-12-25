#include <iostream>
using namespace std;

int main()
{
	//Task("Series21");
	int N;
	cin >> N;
	double number, prev;
	cin >> prev;
	bool progress = true;
	for (int i = 2; i <= N; i++)
	{
		cin >> number;
		if (prev > number)
			progress = false;
		prev = number;
	}
	cout << progress;
	
	return 0;
}
