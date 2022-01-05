package uglyNumber3;

import java.math.BigInteger;
import java.util.Scanner;

class Solution {

    private final static int modulo = 1000000007;
    private static BigInteger howManyUglyOf(BigInteger k, BigInteger a, BigInteger b, BigInteger c) {
        BigInteger acInt = a.multiply(c).divide(a.gcd(c));
        BigInteger abInt = a.multiply(b).divide(a.gcd(b));
        BigInteger bcInt = b.multiply(c).divide(b.gcd(c));
        BigInteger abcInt = abInt.multiply(c).divide(abInt.gcd(c));
        BigInteger uglyCount = (k.divide(a)).add(k.divide(b)).add(k.divide(c))
                .subtract(k.divide(acInt)).subtract(k.divide(abInt)).subtract(k.divide(bcInt))
                .add(k.divide(abcInt));
        return uglyCount;
    }

    public static int nthUglyNumber(int n, int a, int b, int c) {
        if (n == 0 || a == 0 || b == 0 || c == 0) return 0;
        BigInteger aInt = BigInteger.valueOf(a);
        BigInteger bInt = BigInteger.valueOf(b);
        BigInteger cInt = BigInteger.valueOf(c);
        BigInteger maxArg = aInt.max(bInt).max(cInt);
        BigInteger low = BigInteger.valueOf(0);
       BigInteger high = BigInteger.valueOf(n).multiply(BigInteger.valueOf(a).multiply(BigInteger.valueOf(b).multiply(BigInteger.valueOf(c))));
       BigInteger mid = BigInteger.valueOf(0);
       BigInteger k = BigInteger.ZERO;
       while (!k.equals(BigInteger.valueOf(n))) {
           mid = low.add((high.subtract(low).divide(BigInteger.valueOf(2))));
           if (high.subtract(low).compareTo(BigInteger.valueOf(1)) < 1) break;
           k = howManyUglyOf(mid, aInt, bInt, cInt);
           if (k.compareTo(BigInteger.valueOf(n))>0) {
               high = mid.subtract(BigInteger.valueOf(1));
           } else  if (k.compareTo(BigInteger.valueOf(n))<0) {
               low = mid.add(BigInteger.valueOf(1));
           }
       }
       if (mid.compareTo(maxArg)<0) return maxArg.mod(BigInteger.valueOf(modulo)).intValue();
       return mid.mod(BigInteger.valueOf(modulo)).intValue();
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = Integer.parseInt(scanner.nextLine());
        int a = Integer.parseInt(scanner.nextLine());
        int b = Integer.parseInt(scanner.nextLine());
        int c = Integer.parseInt(scanner.nextLine());
        scanner.close();
        int res = nthUglyNumber(n, a, b, c);
        System.out.println(res);
    }
}