#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File55");
	char S0[100], S[100];
	int N;
	cin.getline(S0, 100);
	cin >> N;
	cin.ignore();
	ifstream fileIn;
	ofstream fileOut(S0, ios::binary);
	int number, elements;
	for (int i = 0; i < N; i++)
	{
		cin.getline(S, 100);
		fileIn.open(S, ios::binary);
		fileIn.seekg(0, ios::end);
		elements = fileIn.tellg() / sizeof(int);
		fileIn.seekg(0);
		fileOut.write(reinterpret_cast<const char*>(&elements), sizeof(elements));
		for (int j = 0; j < elements; j++)
		{
			fileIn.read(reinterpret_cast<char*>(&number), sizeof(number));
			fileOut.write(reinterpret_cast<const char*>(&number), sizeof(number));
		}
		fileIn.close();
	}
	fileOut.close();
	
	return 0;
}
