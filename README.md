# Tracking

This tool helps generating invoices from data stored in yaml files.

There are three different file for each kind of data; clients, hours and invoices.

Run `./tracking -h` to see the available options.

## Clients file

The template for the clients file is:

```yaml
---
- company: <CLIENT_KEY>
  name: <CLIENT_NAME>
  projects:
    - key: <PROJECT_KEY>
      name: <PROJECT_NAME>
  billing: <BILLING_INFO>
```

For example:

```yaml
---
- company: foobar
  name: Foo Bar Ltd
  projects:
    - key: p1
      name: Project 1
    - key: p2
      name: Project 2
  billing: |-
    Foo Bar Manager
    Foo Bar Ltd (12345678)
    foo@bar.io
    12 Foo bar st, Foobarland
    VAT: FB12345678
    MOD: 98273311

- company: zap
  billing: |-
    ZAP Manager
    ZAP Ltd (12345678)
    zap@email.io
    12 ZAP st, Foobarland
    VAT: FB87654321
    MOD: 8877665
```

## Hours file

The template for the hours file is:

```yaml
---
- day: <DAY>
  hours:
    <CLIENT_KEY>.<PROJECT_KEY>: <NUM>
```

For example:

```yaml
---
- day: 2018-09-12
  hours:
    foobar.p1: 1
    zap.p3: 5

- day: 2018-09-13
  hours:
    foobar.p2: 1
    zap.p5: 3
```

## Invoices file

The template for the hours file is:
```yaml
---
- invoice: <INVOICE_ID>
  client: <CLIENT_KEY>
  rate: <HOURLY_RATE>
  from: <FROM_DATE>
  to: <TO_DATE>
```

For example:

```yaml
---
- invoice: 100001
  client: foobar
  rate: 12.5
  from: 2018-09-12
  to: 2018-09-13

- invoice: 100002
  client: foobar
  rate: 12.5
  from: 2018-09-14
  to: 2018-09-15
```
