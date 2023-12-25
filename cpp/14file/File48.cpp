#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File48");
	char SA[100], SB[100], SC[100], SD[100];
	cin.getline(SA, 100);
	cin.getline(SB, 100);
	cin.getline(SC, 100);
	cin.getline(SD, 100);
	ifstream fileA(SA, ios::binary);
	ifstream fileB(SB, ios::binary);
	ifstream fileC(SC, ios::binary);
	ofstream fileD(SD, ios::binary);
	int number;
	while (true)
	{
		fileA.read(reinterpret_cast<char*>(&number), sizeof(number));
		if (fileA.eof())
		{
			fileA.clear();
			break;
		}
		fileD.write(reinterpret_cast<const char*>(&number), sizeof(number));
		
		fileB.read(reinterpret_cast<char*>(&number), sizeof(number));
		fileD.write(reinterpret_cast<const char*>(&number), sizeof(number));

		fileC.read(reinterpret_cast<char*>(&number), sizeof(number));
		fileD.write(reinterpret_cast<const char*>(&number), sizeof(number));
	}
	fileA.close();
	fileB.close();
	fileC.close();
	fileD.close();
	
	return 0;
}
