#include <stdio.h>
#include <string.h>

int main(){
  int x, y, i;
  char day[7][4] = {"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"};
  int daybymonth[12] = {1, 32, 60, 91, 121, 152, 182, 213, 244, 274, 305, 335};
 //                     1   2   3   4    5    6    7    8    9   10   11   12
  scanf("%d %d", &x, &y);
  printf("%s",day[(daybymonth[x-1]+(y-1))%7]);

  return 0;

}
