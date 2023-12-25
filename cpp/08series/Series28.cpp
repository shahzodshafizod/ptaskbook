#include <iostream>
using namespace std;

int main()
{
	//Task("Series28");
	int N;
	cin >> N;
	double number, degree;
	for (int i = 1; i <= N; i++)
	{
		cin >> number;
		degree = 1;
		for (int j = i; j <= N; j++)
			degree *= number;
		cout << degree;
	}
	
	return 0;
}
