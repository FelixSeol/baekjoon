package bj_2178;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.Queue;
import java.util.Scanner;

public class Main {
	public static int[] dx={0, 0, 1, -1};
	public static int[] dy={1, -1, 0, 0};
	public static int N;
	public static int M;
	public static int[][] maze;
	public static int[][] adj;
	public static boolean[] visited;
	public static Queue<Point> q;
	public static class Point{
		int x;
		int y;
		int cnt;
		public Point(int x, int y, int cnt){
			this.x = x;
			this.y = y;
			this.cnt = cnt;
		}
	}
	public static void main(String[] args){
		Scanner scanner = new Scanner(System.in);
		N = scanner.nextInt();
		M = scanner.nextInt();
		maze = new int[N][M];
		visited = new boolean[N*M];

		for(int i = 0; i < N; i++){
			String str = scanner.next();
			for(int j = 0; j < M; j++){
				maze[i][j] = (int)str.charAt(j)-48;
			}
		}
		Arrays.fill(visited, false);
		q = new LinkedList<Point>();
		
		
		Point point = new Point(0,0,1);
		visited[0]=true;
		q = new LinkedList<Point>();
		q.offer(point);
		while(true){
			point = q.poll();
			
			for(int i = 0; i < 4; i++){
				if(point.x+dx[i]>=0 && point.x+dx[i]<M && point.y+dy[i] >=0 && point.y+dy[i] <N){
					if(maze[point.y+dy[i]][point.x+dx[i]]==1 && !visited[(point.y+dy[i])*M+point.x+dx[i]]){
						if(point.x+dx[i] == M-1 && point.y+dy[i] == N-1){
							System.out.println(point.cnt+1);
							return;
						}
						visited[(point.y+dy[i])*M + point.x+dx[i]]=true;
						q.offer(new Point(point.x+dx[i], point.y+dy[i], point.cnt+1));
					}
				}
			}
		}
	}
}
