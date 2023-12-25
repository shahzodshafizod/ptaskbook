#include <iostream>
#include <cmath>
using namespace std;

void RectPS(float x1, float y1, float x2, float y2, float& P, float& S);

int main()
{
	//Task("Proc5");
	float x1, y1, x2, y2, P, S;
	for (int i = 1; i <= 3; i++)
	{
		cin >> x1 >> y1 >> x2 >> y2;
		RectPS(x1, y1, x2, y2, P, S);
		cout << P << S;
	}
	
	return 0;
}

void RectPS(float x1, float y1, float x2, float y2, float& P, float& S)
{
	float a = abs(x2 - x1);
	float b = abs(y2 - y1);

	P = 2 * (a+b);
	S = a*b;

}