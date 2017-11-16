package bj_11053;

import java.util.Arrays;
import java.util.Scanner;

public class Main {
	public static int N;
	public static int[] data = null;
	public static int[] score = null;
	public static void main(String[] args){
		Scanner scanner = new Scanner(System.in);
		
		N = Integer.parseInt(scanner.nextLine());
		
		score = new int[N];
		
		data = new int[N];
		for (int i = 0; i < N; i++){
			data[i] = scanner.nextInt();
		}
		
		Arrays.fill(score, 0);
		score[0] = 1;
		int longest = 1;
		for (int i = 1; i < N; i++){
			score[i] = getLongest(i)+1;
		}
		int answer = 0;
		for (int i = 0; i < N; i++){
			if(score[i] > answer)
				answer = score[i];
		}
		System.out.print(answer);
	}
	public static int getLongest(int i){
		int longestScore = 0;
		for(int j = i-1; j >= 0; j--){
			if(data[j] < data[i]){
				if (score[j]>longestScore)
					longestScore = score[j];
			}	
		}
		return longestScore;
	}
}
