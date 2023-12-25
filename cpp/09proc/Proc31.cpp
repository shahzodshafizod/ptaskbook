#include <iostream>
using namespace std;

bool IsPalindrome(int K);

int main()
{
	//Task("Proc31");
	int number, palindromes = 0;
	for (int i = 1; i <= 10; i++)
	{
		cin >> number;
		palindromes += (int)IsPalindrome(number);
	}
	cout << palindromes;
	
	return 0;
}

bool IsPalindrome(int K)
{
	int temp = K, chappa = 0;
	while (temp > 0)
	{
		chappa = chappa * 10 + temp % 10;
		temp /= 10;
	}
	return chappa == K;
}