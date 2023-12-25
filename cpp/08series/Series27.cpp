#include <iostream>
using namespace std;

int main()
{
	//Task("Series27");
	int N;
	cin >> N;
	double number, degree;
	for (int i = 1; i <= N; i++)
	{
		cin >> number;
		degree = 1;
		for (int j = 1; j <= i; j++)
			degree *= number;
		cout << degree;
	}
	
	return 0;
}
