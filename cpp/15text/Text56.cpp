#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text56");
	const int stringSize = 100, charCounts = 256;

	char textFileName[100], charFileName[100];
	cin.getline(textFileName, 100);
	cin.getline(charFileName, 100);

	ifstream textFile(textFileName);
	ofstream charFile(charFileName, ios::binary);

	char str[stringSize+1];
	bool chars[charCounts] = {false};
	int len, code;
	while (textFile.peek() != -1)
	{
		textFile.getline(str, stringSize+1);
		len = strlen(str);
		for (int i = 0; i < len; i++)
		{
			code = int(str[i]);
			code += code < 0 ? 256 : 0;
			chars[code] = true;
		}
	}
	
	textFile.close();
	char C;
	for (int i = charCounts-1; i >= 0; i--)
		if (chars[i])
		{
			C = char(i);
			charFile.write(reinterpret_cast<const char*>(&C), sizeof(C));
		}

	charFile.close();

	return 0;
}
