/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-13
 * @fileoverview This program tells you if a given year is a leap year.
 */

// variables
let yearAsString: string = "";
let yearAsNumber: number = 0;

// input year
yearAsString = prompt("Please enter a year: ") || "0";
yearAsNumber = parseInt(yearAsString);

// check if leap year
if ((yearAsNumber % 4 === 0 && yearAsNumber % 100 !== 0) || (yearAsNumber % 400 === 0)) {
  console.log(yearAsNumber + " is a leap year.");
} else {
  console.log(yearAsNumber + " is not a leap year.");
}

console.log("\nDone.");