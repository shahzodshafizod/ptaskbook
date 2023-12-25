#include <iostream>
#include <fstream>
using namespace std;

int CharToInt(char C);

struct Date
{
	int day;
	int month;
	int year;
	int pos;
};

int main()
{
	//Task("File72");
	const int stringSize = 80;
	char fileName[100], str[stringSize];
	cin.getline(fileName, 100);
	ifstream file(fileName, ios::binary);
	Date date, lateDate;
	lateDate.pos = -1;
	while (true)
	{
		date.pos = file.tellg();
		file.read(reinterpret_cast<char*>(&str), sizeof(char)*stringSize);
		if (file.eof())
		{
			file.clear();
			break;
		}
		date.month = CharToInt(str[3]) * 10 + CharToInt(str[4]);
		switch (date.month)
		{
			case 9: case 10: case 11:
				date.day = CharToInt(str[0]) * 10 + CharToInt(str[1]);
				date.year = 0;
				for (int j = 6; j < 10; j++)
					date.year = date.year * 10 + CharToInt(str[j]);
				if (lateDate.pos < 0)
					lateDate = date;

				else
				{
					if (date.year > lateDate.year)
						lateDate = date;

					else if ( (date.year == lateDate.year) && (date.month > lateDate.month) )
						lateDate = date;

					else if ( (date.year == lateDate.year) && (date.month == lateDate.month) && (date.day > lateDate.day) )
						lateDate = date;
				}
				break;
		}
	}
	str[0] = '\0';
	if (lateDate.pos >= 0)
	{
		file.seekg(lateDate.pos);
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