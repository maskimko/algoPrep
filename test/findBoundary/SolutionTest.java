package findBoundary;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void findBoundary() {
        boolean[] input  = new boolean[]{false,false,true,true,true};
        int expect = 2;
        assertEquals(expect, Solution.findBoundary(input));
    }
    @Test
    void findBoundaryTrue() {
        boolean[] input  = new boolean[]{true,true,true,true,true};
        int expect = 0;
        assertEquals(expect, Solution.findBoundary(input));
    }
}