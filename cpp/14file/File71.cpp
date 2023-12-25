#include <fstream>
#include <iostream>
using namespace std;

struct Date
{
	int day;
	int month;
	int year;
	int pos;
};

int CharToInt(char C);

int main()
{
	//Task("File71");
	const int stringSize = 80;
	char fileName[100], str[stringSize];
	cin.getline(fileName, 100);
	ifstream file(fileName, ios::binary);
	Date date, earlyDate;
	earlyDate.pos = -1;
	while (true)
	{
		date.pos = file.tellg();
		file.read(reinterpret_cast<char*>(&str), sizeof(char)*stringSize);
		if (file.eof())
		{
			file.clear();
			break;
		}
		date.month = CharToInt(str[3])*10 + CharToInt(str[4]);
		switch (date.month)
		{
			case 3: case 4: case 5:
				date.day = CharToInt(str[0]) * 10 + CharToInt(str[1]);
				date.year = 0;
				for (int j = 6; j < 10; j++)
					date.year = date.year * 10 + CharToInt(str[j]);
				
				if (earlyDate.pos < 0)
					earlyDate = date;
				
				else
				{
					if (date.year < earlyDate.year)
						earlyDate = date;

					else if ( (date.year == earlyDate.year) && (date.month < earlyDate.month) )
						earlyDate = date;
					
					else if ( (date.year == earlyDate.year) && (date.month == earlyDate.month) && (date.day < earlyDate.day) )
						earlyDate = date;
				}
				break;
		}
	}
	str[0] = '\0';
	if (earlyDate.pos >= 0)
	{
		file.seekg(earlyDate.pos);
		file.read(reinterpret_cast<char*>(&str), sizeof(char)*stringSize);
	}
	file.close();
	cout << str;
	
	return 0;
}

int CharToInt(char C)
{
	int farq = int(C) - int('0');
	if ( (farq < 0) || (farq > 9) )
		return -1;
	return farq;
}