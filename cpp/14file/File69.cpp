#include <iostream>
#include <fstream>
using namespace std;

int CharToInt(char C);

int main()
{
	//Task("File69");
	const int stringSize = 80;
	char fileName[100], summerFileName[100];
	cin.getline(fileName, 100);
	cin.getline(summerFileName, 100);
	ifstream file(fileName, ios::binary);
	ofstream summerFile(summerFileName, ios::binary);
	char str[stringSize];
	int month;
	while (true)
	{
		file.read(reinterpret_cast<char*>(&str), sizeof(str));
		if (file.eof())
		{
			file.clear();
			break;
		}
		month = CharToInt(str[3])*10 + CharToInt(str[4]);
		switch (month)
		{
			case 6: case 7: case 8:
				summerFile.write(reinterpret_cast<const char*>(&str), sizeof(char)*stringSize);
				break;
		}
	}
	file.close();
	summerFile.close();
	
	return 0;
}

int CharToInt(char C)
{
	int farq = int(C) - int('0');
	if ( (farq < 0) || (farq > 9) )
		return -1;
	return farq;
}