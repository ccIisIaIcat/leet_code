class Solution:
    def isValidSudoku(self, board) :
        for i in range(9):
            for j in range(9):
                if board[i][j] == ".":
                    board[i][j] = 0
                else:
                    board[i][j] = int(board[i][j])
        for i in range(9):
            for j in range(9):
                if board[i][j] != 0:
                    for k in range(9):
                        if k != j and board[i][k] == board[i][j]:
                            return False
                        if k != i and board[k][j] == board[i][j]:
                            return False
                    a_start = int(i/3)*3
                    b_start = int(j/3)*3
                    for a in range(3):
                        for b in range(3):
                            if (a+a_start)!=i and (b+b_start)!=j and board[a+a_start][b+b_start] == board[i][j]:
                                return False
        return True





if __name__ == '__main__':
    print("hello")
    test = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
    S = Solution()
    print(S.isValidSudoku(test))
