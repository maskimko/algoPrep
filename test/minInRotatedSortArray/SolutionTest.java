package minInRotatedSortArray;

import org.junit.jupiter.api.Test;

import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void findMinRotated1() {
        assertEquals(3, Solution.findMinRotated(Arrays.asList(30, 40, 50, 10, 20)));
    }
    @Test
    void findMinRotated2() {
        assertEquals(7, Solution.findMinRotated(Arrays.asList(3, 5, 7, 11, 13, 17, 19, 2)));
    }
}