package validParentheses;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void validParentheses() {
        assertTrue(Solution.validParentheses("()"));
        assertFalse(Solution.validParentheses("(}"));
    }
}