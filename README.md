# WizTools.org Money Lib

This is a golang module for representing money value.

## Uses of Money Lib

1. You have collected money value and currency type as input, and need this validated. Validation is in terms of the number of pennies that make a dollar value. For some currencies, 100 pennies make a dollar, for others it may be 1000, and some others it could be 0.
2. You are loading the money value from DB, and want to store the money representation in memory for the further computation. The `Money` struct and it's methods give various representations of the money value like in `big.Float` or in whole number form (in pennies).
3. You need a human representation of the money value for the UI layer. For example, USD `1234567.89` for **English** and **Italian** languages are represented thus respectively: `$1,234,567.89`, `$1.234.567,89`.

## Foundation

The library is written on the foundation of NOD, or **Number Of Decimals**, for each currency. For example, for USD, 100 cents make a dollar, which means need to represent pennies in 2 digits (00-to-99) before bumping up the dollar part, and hence the NOD is 2. But JPY does not have the concept of a cent, meaning the NOD for that currency is 0.

## Usage

We encourage you to look at the test cases in the code for examples.
