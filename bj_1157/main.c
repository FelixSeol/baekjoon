#include <stdio.h>
#include <string.h>

int main(void) {
	int i=0;
	char word[1000001] = { 0 };
	int arr[26] = { 0, };
	char result = NULL;
	int max = 0;
	int len;
	
	gets(word);
	len = strlen(word);
	for (i = 0; i < len; i++){
		if (word[i] >= 97 && word[i] <= 122) {
			arr[word[i] - 97]++;
		}
		else{
			arr[word[i] - 65]++;
		}
	}

	for (i = 0; i < 26; i++) {
		if (max < arr[i]) {
			max = arr[i];
			result = i + 65;
		}
		else if (max == arr[i]) {
			result = '?';
		}
		else {};
	}
	
	printf("%c", result);
	return 0;
}
