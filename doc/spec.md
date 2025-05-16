# Receipt reports (Kvittorapports)
Receipt reports are the core of FlashCow. How they are used goes something like this:
1. Someone in the chapter buys something with the chapter card for the chapter.
2. They take a picture of the physical receipt on the FlashCow web app. (Legibility of the image must be confirmed)
3. They put in the necessary details.
4. They commit it.
5. Users with the _bookkeeper_ role can then look at all the receipt reports.
6. The bookkeeper then approves the receipt report.

## Receipt report fields
- Picture of the receipt
- User
- Committee
- Date of purchase
- Contents of the purchase
- Comment field
- If it contains alcohol:
    - Partial sums for alcohol
    - Volume estimations for alcohol
- Non-alcoholic partial sums
- Internrep sum

# The database
## Data redundancy
Data redundancy for the database is handled by having a failover read replica, shapshots of which are taken and uploaded to three other locations.
