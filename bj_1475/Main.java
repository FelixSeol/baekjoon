import java.util.Scanner;

public class Main {
	static String N;
	static int[] arr = new int[10];
	static int Answer = 0;
	public static void main(String[] args){
		Scanner scanner = new Scanner(System.in);
		
		N = scanner.next();
		
		for(int i = 0; i < N.length(); i++){
			arr[N.charAt(i)-48]++;
		}
		arr[6] = (int) Math.ceil((double)(arr[6]+arr[9])/2);
		arr[9] = arr[6];
		
		
		for(int i = 0; i < 10; i++){
			Answer = Math.max(Answer, arr[i]);
		}
		System.out.print(Answer);
	}
}

