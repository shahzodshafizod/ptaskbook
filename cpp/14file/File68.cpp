#include <iostream>
#include <fstream>
using namespace std;

int CharToInt(char C);

int main()
{
	//Task("File68");
	const int stringSize = 80;
	char fileName[100], monthsFileName[100], yearsFileName[100];
	cin.getline(fileName, 100);
	cin.getline(monthsFileName, 100);
	cin.getline(yearsFileName, 100);
	ifstream file(fileName, ios::binary);
	ofstream monthsFile(monthsFileName, ios::binary);
	ofstream yearsFile(yearsFileName, ios::binary);
	file.seekg(0, ios::end);
	int strings = file.tellg() / (stringSize*sizeof(char));
	file.seekg(0);
	char str[stringSize];
	int month, year;
	for (int i = strings-1; i >= 0; i--)
	{
		file.seekg(i*stringSize);
		file.read(reinterpret_cast<char*>(&str), sizeof(char)*stringSize);
		month = CharToInt(str[3])*10 + CharToInt(str[4]);
		year = 0;
		for (int j = 6; j < 10; j++)
			year = year * 10 + CharToInt(str[j]);
		monthsFile.write(reinterpret_cast<const char*>(&month), sizeof(month));
		yearsFile.write(reinterpret_cast<const char*>(&year), sizeof(year));
	}
	file.close();
	monthsFile.close();
	yearsFile.close();
	
	return 0;
}

int CharToInt(char C)
{
	int farq = int(C) - int('0');
	if ( (farq < 0) || (farq > 9) )
		return -1;
	return farq;
}