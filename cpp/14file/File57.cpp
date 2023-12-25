#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File57");
	char S1[100], S2[100], S0[100];
	cin.getline(S1, 100);
	cin.getline(S2, 100);
	cin.getline(S0, 100);
	ofstream fileOut1(S1, ios::binary);
	ofstream fileOut2(S2, ios::binary);
	ifstream fileIn(S0, ios::binary);
	int number, elements;
	while (true)
	{
		fileIn.read(reinterpret_cast<char*>(&elements), sizeof(elements));
		if (fileIn.eof())
		{
			fileIn.clear();
			break;
		}
		fileIn.read(reinterpret_cast<char*>(&number), sizeof(number));
		fileOut1.write(reinterpret_cast<const char*>(&number), sizeof(number));
		fileIn.seekg((elements-2)*sizeof(int), ios::cur);
		fileIn.read(reinterpret_cast<char*>(&number), sizeof(number));
		fileOut2.write(reinterpret_cast<const char*>(&number), sizeof(number));
	}
	fileOut1.close();
	fileOut2.close();
	fileIn.close();
	
	return 0;
}
