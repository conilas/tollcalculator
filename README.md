# Toollulator

[![Build Status](https://github.com/conilas/tollcalculator/workflows/test%20and%20build/badge.svg)](https://github.com/conilas/tollcalculator/actions?workflow=test%20and%20build)
[![Coverage Status](https://coveralls.io/repos/github/conilas/tollcalculator/badge.svg?branch=master)](https://coveralls.io/github/conilas/tollcalculator?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/conilas/tollcalculator)](https://goreportcard.com/report/github.com/conilas/tollcalculator)


## Tollcaculator, CLI style!

This project is a simple solution for the challenge proposed. Above one can see the code coverage, build status and such - those were all done using Github Actions.

The porject uses `gomod`, so one needs go at version 11 or plus. If such is installed, just run:

```
git clone git@github.com:conilas/tollcalculator.git
cd tollcalculator
go build
```

And you should have a build, ready to roll project!

## How to use it

It is, as written, a CLI. So in order to invoke, assuming the build binary from the aforementioned steps is `tollcalculator`, one just run:

```
./tollcalculator -vehicle $VEHICLE_TYPE -timestamps $COMMA_SEPARATED_TIMESTAMPS
```

Where: 

- **$VEHICLE_TYPE** stands for a type of the vehicle. This values should be in a range from 0 to 7, and each one of them represents one type of vehicle. Those are (in order, starting from 0): **Car**, **Truck**, **Motorbike**, **Tractor**, **Emergency**, **Diplomat**, **Foreign**, **Military**. If another value is passed, one may expected the CLI to log an error and exit.
- **$COMMA_SEPARATED_TIMESTAMPS** stands for, as you'd expect :smile:, comma separated UNIX timestamps. Those timestamps will be read as Stockholm (GMT +2) times (i.e **1585839171** stands for **2020-04-02 16:52:51**), so be aware of that.

One concrete example:

```
 ./tollcalculator -vehicle 0 -timestamps 1585839171
```

This will calculate a toll at the time **2020-04-02 16:52:51** for a **Car**. This means that it will log the toll as:

```
Collected tolls are [18]
```

Which is a High price for that moment. (:
