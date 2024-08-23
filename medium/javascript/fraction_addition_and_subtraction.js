// link to the problem: https://leetcode.com/problems/fraction-addition-and-subtraction/

/**
 * @param {string} expression
 * @return {string}
 */
var fractionAddition = function (expression) {
  function gcd(a, b) {
    return b === 0 ? a : gcd(b, a % b);
  }

  const fractions = expression.match(/[+-]?\d+\/\d+/g);

  let numerator = 0;
  let denominator = 1;

  for (const frac of fractions) {
    const [n, d] = frac.split("/").map(Number);

    numerator = numerator * d + denominator * n;
    denominator = denominator * d;

    const commonDivisor = gcd(Math.abs(numerator), denominator);

    numerator /= commonDivisor;
    denominator /= commonDivisor;
  }

  if (denominator < 0) {
    denominator = -denominator;
    numerator = -numerator;
  }

  if (numerator === 0) {
    return "0/1";
  }

  return `${numerator}/${denominator}`;
};
