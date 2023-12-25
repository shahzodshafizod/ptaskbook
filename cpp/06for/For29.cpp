#include <iostream>
using namespace std;

int main()
{
	//Task("For29");
	int N;
	cin >> N;
	double A, B;
	cin >> A >> B;
	double H = (B - A) / N;
	cout << H;
	for (int i = 0; i <= N; i++)
		cout << A+i*H;
	
	return 0;
}
