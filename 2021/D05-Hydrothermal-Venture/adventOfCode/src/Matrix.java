public class Matrix {
    private int[][] board;
    
    public Matrix(int dimension) {
        board = new int[dimension][dimension];
    }

    public void draw(Line line) {
        Coordinate from = line.getFirstCoordinate();
        Coordinate to = line.getSecondCoordinate();

        boolean horizontal = 
        (from.x() == to.x() &&
        from.y() < to.y());

        boolean vertical = 
        (from.x() < to.x() &&
        from.y() == to.y());
        if (from.x() < to.x()) {
            // from left to right
            // horizontal
            if (from.y() == to.y()) {
                for (int i = from.x(); i <= to.x(); i++) {
                    board[i][from.y()]++;
                }
            }

            if (!isDiagonal45Degrees(from, to)) {
                return;
            }

            // diagonal
            int round = 0;
            for (int i = from.x(); i <= to.x(); i++) {
                if (from.y() < to.y()) {
                    // from up to down
                    board[i][from.y() + round]++;
                } else {
                    // from down to up
                    board[i][from.y() - round]++;
                }
                round++;
            }
        } else if (from.x() > to.x()) {
            // from right to left
            // horizontal
            if (from.y() == to.y()) {
                for (int i = from.x(); i >= to.x(); i--) {
                    board[i][from.y()]++;
                }
            }

            if (!isDiagonal45Degrees(from, to)) {
                return;
            }

            // diagonal
            int round = 0;
            for (int i = from.x(); i >= to.x(); i--) {
                if (from.y() < to.y()) {
                    // from up to down
                    board[i][from.y() + round]++;
                } else {
                    // from down to up
                    board[i][from.y() - round]++;
                }
                round++;
            }
        } else {
            // vertical
            if (from.y() < to.y()) {
                for (int j = from.y(); j <= to.y(); j++) {
                    // from up to down
                    board[from.x()][j]++;
                }
            } else if (from.y() > to.y()) {
                for (int j = from.y(); j >= to.y(); j--) {
                    // from down to up
                    board[from.x()][j]++;
                }
            }
        
        }
    }

    public int getOverlapCount() {
        int overlapCount = 0;

        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[i].length; j++) {
                if (board[i][j] >= 2) {
                    overlapCount++;
                }
            }
        }

        return overlapCount;
    }
    
    private boolean isDiagonal45Degrees(final Coordinate from, final Coordinate to) {
        return Math.abs(from.x() - to.x()) == Math.abs(from.y() - to.y());
    }

    public String toString() {
        String matrixStr = "";

        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[i].length; j++) {
                matrixStr += board[i][j];
                if (j != board[i].length - 1) {
                    matrixStr += "\t";    
                }
            }
            if (i != board.length -1) {
                matrixStr += "\n\n";
            }
        }

        return matrixStr;
    }
}
