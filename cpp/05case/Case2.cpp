#include <iostream>
using namespace std;

int main()
{
	//Task("Case2");
	int K;
	cin >> K;
	switch (K)
	{
		case 1: cout << "плохо"; break;
		case 2: cout << "неудовлетворительно"; break;
		case 3: cout << "удовлетворительно"; break;
		case 4: cout << "хорошо"; break;
		case 5: cout << "отлично"; break;
		default: cout << "ошибка"; break;
	}
	
	return 0;
}
