package bj_3040;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.Queue;
import java.util.Scanner;

public class Main {
	public static int[] arr;
	public static int sum = 0;
	
	public static void main(String[] args){
		Scanner scanner = new Scanner(System.in);
		arr = new int[9];
	
		for(int i = 0; i < 9; i++){
			arr[i] = Integer.parseInt(scanner.nextLine());
		}
	
		Solve(0, 0, "");
	}
	
	public static void Solve(int N, int cnt, String answer){
		String temp = answer.toString();
		
		if(N == 9){
			return;
		}
		if((sum + arr[N] == 100) && (cnt==6)){
			answer = answer+arr[N];
			System.out.println(answer);
			return;
		}
		Solve(N+1, cnt, temp);
		
		sum += arr[N];
		Solve(N+1, cnt+1, temp+arr[N]+"\n");
		
		sum -= arr[N];
	}
}
