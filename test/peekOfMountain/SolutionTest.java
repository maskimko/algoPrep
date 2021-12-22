package peekOfMountain;

import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void peakOfMountainArray() {
        assertEquals(3,Solution.peakOfMountainArray(Arrays.asList(0,1,2,3,2,1,0)));
        assertEquals(5,Solution.peakOfMountainArray(Arrays.asList(0,1,2,3,4,5,4,2,0)));
    }

    @Test
    void peakOfMountainArray2() {
        assertEquals(1,Solution.peakOfMountainArray(Arrays.asList(0,10,3,2,1)));
    }


    @Test
    void peakOfMountainArray3() {
        assertEquals(1,Solution.peakOfMountainArray(Arrays.asList(0,10,0)));
    }
}