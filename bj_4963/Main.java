package bj_4963;

import java.util.Arrays;
import java.util.Scanner;

public class Main {
	public static int w = 1;
	public static int h = 1;
	public static int Answer = 0;
	public static int[][] map = null;
	public static int[] union = null; // union[x][y] : x's root is y
	public static void main(String[] args){
		Scanner scanner = new Scanner(System.in);
		w = scanner.nextInt();
		h = scanner.nextInt();
		while(w!=0 && h!=0){
			Answer = 0;
			
			//System.out.println("width, height setted");
			map = new int[h][w];
			union = new int[h*w];
			for(int i = 0; i < h; i++){
				for(int j = 0; j < w; j++){
					map[i][j] = scanner.nextInt();
				}
				scanner.nextLine();
			}
			//System.out.println("map setted");
			Arrays.fill(union, -1);
			for(int i = 0; i < h; i++){
				for(int j = 0; j < w; j++){
					if(map[i][j]==1){
						//if map[i][j] == 1 then my root is myself
						if(i==0&&j==0){
							
						}else if(i == 0){
							if(map[i][j-1]==1){
								Union(i*w+j, i*w+j-1);
							}
						}else if (j == 0){
							if(map[i-1][j]==1){
								Union(i*w+j, (i-1)*w+j);
							}
							if(map[i-1][j+1]==1){
								Union(i*w+j, (i-1)*w+j+1);
							}
						}else if (j == w-1){
							if(map[i][j-1]==1){
								Union(i*w+j, i*w+j-1);
							}
							if(map[i-1][j]==1){
								Union(i*w+j, (i-1)*w+j);
							}
							if(map[i-1][j-1]==1){
								Union(i*w+j, (i-1)*w+j-1);
							}
						}else{
							if(map[i][j-1]==1){
								Union(i*w+j, i*w+j-1);
							}
							if(map[i-1][j]==1){
								Union(i*w+j, (i-1)*w+j);
							}
							if(map[i-1][j-1]==1){
								Union(i*w+j, (i-1)*w+j-1);
							}
							if(map[i-1][j+1]==1){
								Union(i*w+j, (i-1)*w+j+1);
							}
						}
					}else{
						union[i*w+j]=-2;
					}
				}
			}
			
			/*print
			for(int i = 0; i < h; i++){
				for(int j = 0; j < w; j++){
					System.out.printf("%d\t",map[i][j]);	
					}
				System.out.println();
			}
			System.out.println("Unionset");
			for(int i = 0; i < h; i++){
				for(int j = 0; j < w; j++){
					System.out.printf("%d\t",union[i*w+j]);	
					}
				System.out.println();
			}*/
			for (int i = 0 ; i < w*h; i++){
				if(union[i]==-1)
						Answer++;
			}
			System.out.println(Answer);
			w = scanner.nextInt();
			h = scanner.nextInt();
		}			
	}

	public static int Find(int n){
		if(union[n] < 0) return n;
		union[n] = Find(union[n]);
		return union[n];
	}
	
	public static void Union(int a, int b){
		a = Find(a);
		b = Find(b);
		if( a == b ) return;
		union[a] = b;
	}
}
