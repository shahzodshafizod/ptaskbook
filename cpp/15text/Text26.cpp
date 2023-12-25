#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text26");
	const int stringSize = 80;
	char fileName[100], str[stringSize];
	cin.getline(fileName, 100);
	ifstream file(fileName);
	int sarxats = 0;
	while (file.peek() != -1)
	{
		file.getline(str, stringSize);
		if (strlen(str) >= 5)
		{
			bool isSarxat = true;
			for (int i = 0; i < 5; i++)
				if (str[i] != ' ')
					isSarxat = false;
			sarxats += (int)isSarxat;
		}
	}
	file.close();
	cout << sarxats;
	
	return 0;
}
