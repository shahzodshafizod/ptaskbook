#include <iostream>
#include <fstream>
using namespace std;

int CharToInt(char C);

int main()
{
	//Task("File67");
	const int stringSize = 80;
	char fileName[100], daysFileName[100], monthsFileName[100];
	cin.getline(fileName, 100);
	cin.getline(daysFileName, 100);
	cin.getline(monthsFileName, 100);
	ifstream ifs(fileName, ios::binary);
	ofstream daysFile(daysFileName, ios::binary);
	ofstream monthsFile(monthsFileName, ios::binary);
	char str[stringSize];
	int day, month;
	while (true)
	{
		ifs.read(reinterpret_cast<char*>(&str), stringSize*sizeof(char));
		if (ifs.eof())
		{
			ifs.clear();
			break;
		}
		day = CharToInt(str[0])*10 + CharToInt(str[1]);
		month = CharToInt(str[3])*10 + CharToInt(str[4]);
		daysFile.write(reinterpret_cast<const char*>(&day), sizeof(day));
		monthsFile.write(reinterpret_cast<const char*>(&month), sizeof(month));
	}
	ifs.close();
	daysFile.close();
	monthsFile.close();
	
	return 0;
}

int CharToInt(char C)
{
	int farq = int(C) - int('0');
	if ( (farq < 0) || (farq > 9) )
		return -1;
	return farq;
}