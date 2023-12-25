#include <iostream>
#include <fstream>
using namespace std;

struct Date
{
	int day;
	int month;
	int year;
	int pos;
};

int CharToInt(char C);
int datesCompare(const Date& date1, const Date& date2);

int main()
{
	//Task("File73");
	const int stringSize = 80;
	char fileName[100], decFileName[100], str[stringSize];
	cin.getline(fileName, 100);
	cin.getline(decFileName, 100);
	ifstream file(fileName, ios::binary);
	ofstream decFile(decFileName, ios::binary);
	file.seekg(0, ios::end);
	int strings = file.tellg() / (sizeof(char)*stringSize);
	file.seekg(0);
	Date* dates = new Date [strings];
	for (int i = 0; i < strings; i++)
	{
		dates[i].pos = file.tellg();
		file.read(reinterpret_cast<char*>(&str), sizeof(char)*stringSize);
		dates[i].day = CharToInt(str[0]) * 10 + CharToInt(str[1]);
		dates[i].month = CharToInt(str[3]) * 10 + CharToInt(str[4]);
		dates[i].year = 0;
		for (int j = 6; j < 10; j++)
			dates[i].year = dates[i].year * 10 + CharToInt(str[j]);
	}
	Date temp;
	for (int i = 0; i < strings-1; i++)
		for (int j = 1; j < strings-i; j++)
			if (datesCompare(dates[j-1], dates[j]) == -1)
			{
				temp = dates[j];
				dates[j] = dates[j-1];
				dates[j-1] = temp;
			}

	for (int i = 0; i < strings; i++)
	{
		file.seekg(dates[i].pos);
		file.read(reinterpret_cast<char*>(&str), sizeof(char)*stringSize);
		decFile.write(reinterpret_cast<const char*>(&str), sizeof(char)*stringSize);
	}
	
	delete [] dates;
	dates = NULL;
	file.close();
	decFile.close();
	
	return 0;
}

int datesCompare(const Date& date1, const Date& date2)
{
	if (date1.year > date2.year)
		return 1;
	
	if ( (date1.year == date2.year) && (date1.month > date2.month) )
		return 1;

	if ( (date1.year == date2.year) && (date1.month == date2.month) && (date1.day > date2.day) )
		return 1;

	if ( (date1.year == date2.year) && (date1.month == date2.month) && (date1.day == date2.day) )
		return 0;

	return -1;
}

int CharToInt(char C)
{
	int farq = int(C) - int('0');
	if ( (farq < 0) || (farq > 9) )
		return -1;
	return farq;
}