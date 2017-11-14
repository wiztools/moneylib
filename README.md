# WizTools.org Money Lib

This is a golang library for representing money value in a golang application.

## Setup

The library depends on a JSON configuration file that is available in this repository `conf/currency.json`. Any application that depends on this library must have this configuration in the same path.

## Foundation

The library is written on the foundation of NOD, or **Number Of Decimals**, for each currency. For example, for USD, 100 cents make a dollar, and hence the NOD is 2. But JPY does not have the concept of a cent, meaning the NOD for that currency is 0.

## Usage

We encourage you to look at the test cases in the code for examples.
