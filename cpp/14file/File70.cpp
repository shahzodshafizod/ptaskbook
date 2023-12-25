#include <iostream>
#include <fstream>
using namespace std;

int CharToInt(char C);

int main()
{
	//Task("File70");
	const int stringSize = 80;
	char fileName[100], winterFileName[100], str[stringSize];
	cin.getline(fileName, 100);
	cin.getline(winterFileName, 100);
	ifstream file(fileName, ios::binary);
	ofstream winterFile(winterFileName, ios::binary);
	file.seekg(0, ios::end);
	int strings = file.tellg() / (stringSize*sizeof(char));
	int month;
	for (int i = strings-1; i >= 0; i--)
	{
		file.seekg(i*stringSize);
		file.read(reinterpret_cast<char*>(&str), sizeof(char)*stringSize);
		if (file.eof())
		{
			file.clear();
			break;
		}
		month = CharToInt(str[3]) * 10 + CharToInt(str[4]);
		switch (month)
		{
			case 12: case 1: case 2:
				winterFile.write(reinterpret_cast<const char*>(&str), sizeof(char)*stringSize);
				break;
		}
	}
	file.close();
	winterFile.close();
	
	return 0;
}

int CharToInt(char C)
{
	int farq = int(C) - int('0');
	if ( (farq < 0) || (farq > 9) )
		return -1;
	return farq;
}