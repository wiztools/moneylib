# Money Value Storage and Representation

By Subhash Chandran

----

## 1. Introduction

Money is a complicated subject. Different currencies represent money in different forms:

1. 100 pennies make a Dollar. Pennies are represented in 2 decimal precision points.
2. Japanese Yen does not have the concept of pennies. It is always represented without decimal points.
3. 1000 baisa makes one Omani Rial. So Omani currency is represented in 3 decimal precision points.

Considering these differences, financial systems must be flexible enough to store and retrieve money value in all the available currencies in the world in a efficient and accurate way. The objective of this design document is to achieve this.

## 2. Standards

ISO 4217 defines various details of major global currencies including the decimal precision of each of these currencies. Details here:

https://en.wikipedia.org/wiki/ISO_4217

Financial systems need to be designed to support this standard. We are using the ISO 4217 data source from:

https://www.currency-iso.org/en/home/tables/table-a1.html


## 3. Options

Following options are available:

1. Store money value in decimal format in system DB.
2. Store money value in pennies as whole number.
3. Store the integer part and decimal part separately.

The factors that will influence our decision are:

1. FinTech companies work with partner APIs (from banks / accounting partners / etc.).
2. Displaying money for human users.
3. Computing any monetary value charges like fees or commissions.
4. Computing aggregates like total spend over a month.
5. Doing big data analytics.
6. DB design implications.

Let’s discuss the options in detail.

### 3.1. Option 1: Store money value in decimal format

| Parameter | Evaluation |
| --------- | ---------- |
| Partner APIs | TBD |
| Human display | Experience APIs can share the value from DB directly to the frontend systems, where it is rendered with minimal manipulation (adding currency symbol, adding commas per thousand, etc.). |
| Money value computation | Fees and commission computation may become difficult due to differences in currencies, and how we decide to treat the precision. |
| Money value aggregation | Aggregation logic will also become little complicated and need to be written specific to each currency. |
| Big data analytics | TBD |
| DB design | Will require one column to store money value. |

### 3.2. Option 2: Store money value in pennies

| Parameter | Evaluation |
| --------- | ---------- |
| Partner APIs | TBD |
| Human display | Experience APIs need to convert the pennies to decimal format and send to the frontend systems for rendering. |
| Money value computation | Computations are universal, no specific details need to be handled. |
| Money value aggregation | Computations are universal, no specific details need to be handled. |
| Big data analytics | TBD |
| DB design | Will require one column to store money value. |

### 3.3. Option 3: Store integer and decimal part separately

| Parameter | Evaluation |
| --------- | ---------- |
| Partner APIs | TBD |
| Human display | Experience APIs need to combine them and send it across to the frontend system. |
| Money value computation | Needs to be combined, and then the calculation needs to be done. |
| Money value aggregation | Aggregation will require each amount to be combined independently and then added. |
| Big data analytics | TBD |
| DB design | Will require two columns to store money value. |

## 4. Evaluation and selection

As each option has it’s own pros and cons, the evaluation depends on the weightage each evaluation-criteria has. Notes:

1. Top criteria is fees/charges computation and aggregation. Now if the system DB allowed decimal storage of money value, a programmer might store a wrong value. Example: a 5% charge on 1234.56 turns out to be 61.728. Now we are adding a precision value of 0.008 to a currency that may not support that. We expect the programmer to enforce it. It would be better if the integrity is enforced at the DB level. From this perspective, storing money value as a whole number is beneficial. Whole numbers also help in reducing if-else checks in code if some computation is same across multiple currencies.
2. Display money value for human consumption processing is not a high priority, as a flexible system may be rendered in different languages in the future, and each language may have a different way of representing their currency. So this is eventually inevitable.
3. Partner APIs. Typically, any partner API must be considered just one implementation among the many available. The system should be flexible to accommodate any type of representation of money from the partner APIs. Do not design a solution for just one partner API.
4. DB design. This is not a big deal, and it is a one time effort.

Based on these facts, it is decided to store money as a whole number. The Number-of-Decimals, or NOD, will be taken from the currency metadata.
