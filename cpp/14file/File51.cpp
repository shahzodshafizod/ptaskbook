#include <iostream>
#include <fstream>
using namespace std;

void ZipDec(ifstream& from1, ifstream& from2, ofstream& to, double a, double b);

int main()
{
	//Task("File51");
	char SA[100], SB[100], SC[100], SD[100];
	cin.getline(SA, 100);
	cin.getline(SB, 100);
	cin.getline(SC, 100);
	cin.getline(SD, 100);
	ifstream file1(SA, ios::binary);
	ifstream file2(SB, ios::binary);
	ifstream file3(SC, ios::binary);
	ofstream file4(SD, ios::binary);

	double a, b, c;
	file1.read(reinterpret_cast<char*>(&a), sizeof(a));
	file2.read(reinterpret_cast<char*>(&b), sizeof(b));
	file3.read(reinterpret_cast<char*>(&c), sizeof(c));

	while ( !file1.eof() || !file2.eof() || !file3.eof() )
	{
		if (file1.eof())
		{
			file1.clear();
			ZipDec(file2, file3, file4, b, c);
			break;
		}
		if (file2.eof())
		{
			file2.clear();
			ZipDec(file1, file3, file4, a, c);
			break;
		}
		if (file3.eof())
		{
			file3.clear();
			ZipDec(file1, file2, file4, a, b);
			break;
		}

		if ( (a > b) && (a > c) )
		{
			file4.write(reinterpret_cast<const char*>(&a), sizeof(a));
			file1.read(reinterpret_cast<char*>(&a), sizeof(a));
		}
		else if (b > c)
		{
			file4.write(reinterpret_cast<const char*>(&b), sizeof(b));
			file2.read(reinterpret_cast<char*>(&b), sizeof(b));
		}
		else
		{
			file4.write(reinterpret_cast<const char*>(&c), sizeof(c));
			file3.read(reinterpret_cast<char*>(&c), sizeof(c));
		}
	}

	file1.close();
	file2.close();
	file3.close();
	file4.close();
	
	return 0;
}

void ZipDec(ifstream& from1, ifstream& from2, ofstream& to, double a, double b)
{
	while (!from1.eof() || !from2.eof())
	{
		if (from1.eof())
		{
			from1.clear();
			to.write(reinterpret_cast<const char*>(&b), sizeof(b));
			while (true)
			{
				from2.read(reinterpret_cast<char*>(&b), sizeof(b));
				if (from2.eof())
				{
					from2.clear();
					break;
				}
				to.write(reinterpret_cast<const char*>(&b), sizeof(b));
			}
			break;
		}
		
		if (from2.eof())
		{
			from2.clear();
			to.write(reinterpret_cast<const char*>(&a), sizeof(a));
			while (true)
			{
				from1.read(reinterpret_cast<char*>(&a), sizeof(a));
				if (from1.eof())
				{
					from1.clear();
					break;
				}
				to.write(reinterpret_cast<const char*>(&a), sizeof(a));
			}
			break;
		}

		if (a > b)
		{
			to.write(reinterpret_cast<const char*>(&a), sizeof(a));
			from1.read(reinterpret_cast<char*>(&a), sizeof(a));
		}
		else
		{
			to.write(reinterpret_cast<const char*>(&b), sizeof(b));
			from2.read(reinterpret_cast<char*>(&b), sizeof(b));
		}
	}
}