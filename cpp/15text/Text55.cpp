#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text55");
	const int stringSize = 100, charsCounts = 256;
	
	char textFileName[100], charFileName[100];
	cin.getline(textFileName, 100);
	cin.getline(charFileName, 100);

	ifstream textFile(textFileName);
	ofstream charFile(charFileName, ios::binary);

	char str[stringSize+1];
	int len, code;
	
	bool chars[charsCounts] = {false};
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
	char  C;
	for (int i = 0; i < charsCounts; i++)
		if (chars[i])
		{
			C = char(i);
			charFile.write(reinterpret_cast<const char*>(&C), sizeof(C));
		}
	
	charFile.close();

	return 0;
}
