public class Coordinate {
    private final int x, y;

    public Coordinate (int x, int y) {
        this.x = y;
        this.y = x;
    }

    public int x() {
        return this.x;
    }

    public int y() {
        return this.y;
    }

    public String toString() {
        return String.format("(%d,%d)", x, y);
    }
}
