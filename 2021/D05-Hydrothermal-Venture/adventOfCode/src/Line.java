public class Line {
    private final Coordinate p0, p1;

    public Line(Coordinate p0, Coordinate p1) {
        if (p0.x() > p1.x() || p0.y() > p1.y()) {
            this.p0 = p1;
            this.p1 = p0;    
        } else {
            this.p0 = p0;
            this.p1 = p1;
        }
    }

    public Coordinate getFirstCoordinate() {
        return this.p0;
    }

    public Coordinate getSecondCoordinate() {
        return this.p1;
    }

    public String toString() {
        return String.format("%s ---> %s", p0, p1);
    }
}
