package bj_10250;

import java.util.Scanner;
import java.util.StringTokenizer;

public class Main {
	public static int T = 0;
	public static int H = 0;
	public static int W = 0;
	public static int N = 0;
	public static int answerH = 0;
	public static int answerW = 0;
	public static void main(String[] args){
		Scanner scanner = new Scanner(System.in);
		
		T = Integer.parseInt(scanner.nextLine());
		
		for(int i = 0; i < T; i++){
			String str = scanner.nextLine();
			StringTokenizer token = new StringTokenizer(str);
			H = Integer.parseInt(token.nextToken());
			W = Integer.parseInt(token.nextToken());
			N = Integer.parseInt(token.nextToken());
			if(N%H == 0){
				answerH = H;
				answerW = N/H;
			}else{
				answerH = N%H;
				answerW = N/H + 1;
			}
			
			System.out.printf("%d%02d",answerH,answerW);
			
			answerH = 0;
			answerW = 0;
		}
	}
}
